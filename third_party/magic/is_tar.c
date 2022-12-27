#include "file.h"

#ifndef lint
FILE_RCSID("@(#)$File: is_tar.c,v 1.50 2022/12/26 17:31:14 christos Exp $")
#endif

#include "magic.h"
#include <string.h>
#include <ctype.h>
#include "tar.h"

#define	isodigit(c)	( ((c) >= '0') && ((c) <= '7') )

file_private int is_tar(const unsigned char *, size_t);
file_private int from_oct(const char *, size_t);	/* Decode octal number */

static const char tartype[][32] = {	/* should be equal to messages */
	"tar archive",			/* found in ../magic/Magdir/archive */
	"POSIX tar archive",
	"POSIX tar archive (GNU)",	/*  */
};

file_protected int
file_is_tar(struct magic_set *ms, const struct buffer *b)
{
	const unsigned char *buf = CAST(const unsigned char *, b->fbuf);
	size_t nbytes = b->flen;
	/*
	 * Do the tar test first, because if the first file in the tar
	 * archive starts with a dot, we can confuse it with an nroff file.
	 */
	int tar;
	int mime = ms->flags & MAGIC_MIME;

	if ((ms->flags & (MAGIC_APPLE|MAGIC_EXTENSION)) != 0)
		return 0;

	tar = is_tar(buf, nbytes);
	if (tar < 1 || tar > 3)
		return 0;

	if (mime == MAGIC_MIME_ENCODING)
		return 1;

	if (file_printf(ms, "%s", mime ? "application/x-tar" :
	    tartype[tar - 1]) == -1)
		return -1;

	return 1;
}

/*
 * Return
 *	0 if the checksum is bad (i.e., probably not a tar archive),
 *	1 for old UNIX tar file,
 *	2 for Unix Std (POSIX) tar file,
 *	3 for GNU tar file.
 */
file_private int
is_tar(const unsigned char *buf, size_t nbytes)
{
	static const char gpkg_match[] = "/gpkg-1";

	const union record *header = RCAST(const union record *,
	    RCAST(const void *, buf));
	size_t i;
	int sum, recsum;
	const unsigned char *p, *ep;
	const char *nulp;

	if (nbytes < sizeof(*header))
		return 0;

	/* If the file looks like Gentoo GLEP 78 binary package (GPKG),
	 * don't waste time on further checks and fall back to magic rules.
	 */
	nulp = CAST(const char *,
	    memchr(header->header.name, 0, sizeof(header->header.name)));
	if (nulp != NULL && nulp >= header->header.name + sizeof(gpkg_match) &&
	    memcmp(nulp - sizeof(gpkg_match) + 1, gpkg_match,
	    sizeof(gpkg_match)) == 0)
	    return 0;

	recsum = from_oct(header->header.chksum, sizeof(header->header.chksum));

	sum = 0;
	p = header->charptr;
	ep = header->charptr + sizeof(*header);
	while (p < ep)
		sum += *p++;

	/* Adjust checksum to count the "chksum" field as blanks. */
	for (i = 0; i < sizeof(header->header.chksum); i++)
		sum -= header->header.chksum[i];
	sum += ' ' * sizeof(header->header.chksum);

	if (sum != recsum)
		return 0;	/* Not a tar archive */

	if (strncmp(header->header.magic, GNUTMAGIC,
	    sizeof(header->header.magic)) == 0)
		return 3;		/* GNU Unix Standard tar archive */

	if (strncmp(header->header.magic, TMAGIC,
	    sizeof(header->header.magic)) == 0)
		return 2;		/* Unix Standard tar archive */

	return 1;			/* Old fashioned tar archive */
}


/*
 * Quick and dirty octal conversion.
 *
 * Result is -1 if the field is invalid (all blank, or non-octal).
 */
file_private int
from_oct(const char *where, size_t digs)
{
	int	value;

	if (digs == 0)
		return -1;

	while (isspace(CAST(unsigned char, *where))) {	/* Skip spaces */
		where++;
		if (digs-- == 0)
			return -1;		/* All blank field */
	}
	value = 0;
	while (digs > 0 && isodigit(*where)) {	/* Scan til non-octal */
		value = (value << 3) | (*where++ - '0');
		digs--;
	}

	if (digs > 0 && *where && !isspace(CAST(unsigned char, *where)))
		return -1;			/* Ended on non-(space/NUL) */

	return value;
}
