package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"test-station/pkg/rabbitmq"
	"time"
)

//生产者发布流程
func main() {
	var (
		mg sync.WaitGroup
		e = make(chan os.Signal)
	)
	// 初始化mq
	mq := rabbitmq.NewRabbitMQ("queue_publisher", "exchange_publisher", "topsec")
	defer mq.ReleaseRes() // 完成任务释放资源

	// 1.声明队列
	/*
	  如果只有一方声明队列，可能会导致下面的情况：
	   a)消费者是无法订阅或者获取不存在的MessageQueue中信息
	   b)消息被Exchange接受以后，如果没有匹配的Queue，则会被丢弃

	  为了避免上面的问题，所以最好选择两方一起声明
	  ps:如果客户端尝试建立一个已经存在的消息队列，Rabbit MQ不会做任何事情，并返回客户端建立成功的
	*/
	_, err := mq.Channel.QueueDeclare( // 返回的队列对象内部记录了队列的一些信息，这里没什么用
		mq.QueueName, // 队列名
		true,         // 是否持久化
		false,        // 是否自动删除(前提是至少有一个消费者连接到这个队列，之后所有与这个队列连接的消费者都断开时，才会自动删除。注意：生产者客户端创建这个队列，或者没有消费者客户端与这个队列连接时，都不会自动删除这个队列)
		false,        // 是否为排他队列（排他的队列仅对“首次”声明的conn可见[一个conn中的其他channel也能访问该队列]，conn结束后队列删除）
		true,        // 是否阻塞
		nil,          //额外属性（我还不会用）
	)
	if err != nil {
		fmt.Println("声明队列失败", err)
		return
	}

	// 2.声明交换器
	err = mq.Channel.ExchangeDeclare(
		mq.Exchange, //交换器名
		"topic",     //exchange type：一般用fanout、direct、topic
		true,        // 是否持久化
		false,       //是否自动删除（自动删除的前提是至少有一个队列或者交换器与这和交换器绑定，之后所有与这个交换器绑定的队列或者交换器都与此解绑）
		false,       //设置是否内置的。true表示是内置的交换器，客户端程序无法直接发送消息到这个交换器中，只能通过交换器路由到交换器这种方式
		false,       // 是否阻塞
		nil,         // 额外属性
	)
	if err != nil {
		fmt.Println("声明交换器失败", err)
		return
	}

	// 3.建立Binding(可随心所欲建立多个绑定关系)
	err = mq.Channel.QueueBind(
		mq.QueueName,  // 绑定的队列名称
		mq.RoutingKey, // bindkey 用于消息路由分发的key
		mq.Exchange,   // 绑定的exchange名
		false,         // 是否阻塞
		nil,           // 额外属性
	)
	// err = mq.Channel.QueueBind(
	//  mq.QueueName,  // 绑定的队列名称
	//  "routingkey2", // bindkey 用于消息路由分发的key
	//  mq.Exchange,   // 绑定的exchange名
	//  false,         // 是否阻塞
	//  nil,           // 额外属性
	// )
	if err != nil {
		fmt.Println("绑定队列和交换器失败", err)
		return
	}
	mg.Add(1)

	//监听指定信号
	signal.Notify(e, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)

	go func() {
		for  {
			select {
			case <- e:
				fmt.Println("get a signal...")
				mg.Done()
			default:
				time.Sleep(time.Second*10)
				fmt.Println("send a message...")
				// 4.发送消息
				_ = mq.Channel.Publish(
					mq.Exchange,   // 交换器名
					mq.RoutingKey, // routing key
					false,         // 是否返回消息(匹配队列)，如果为true, 会根据binding规则匹配queue，如未匹配queue，则把发送的消息返回给发送者
					false,         // 是否返回消息(匹配消费者)，如果为true, 消息发送到queue后发现没有绑定消费者，则把发送的消息返回给发送者
					amqp.Publishing{ // 发送的消息，固定有消息体和一些额外的消息头，包中提供了封装对象
						ContentType: "text/plain",           // 消息内容的类型
						Body:        []byte("hello jochen " + time.Now().Format("2006-01-02 15:04:05")), // 消息内容
					},
				)
			}

		}
	}()

	mg.Wait()
}