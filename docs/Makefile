# guard against environment variables
MAN1_TXT = $(wildcard deis-accounts-*.txt)
MAN1_TXT += deis-accounts.txt

MAN_TXT = $(MAN1_TXT)
DOC_HTML = $(patsubst %.txt,%.html,$(MAN_TXT))
DOC_MAN1 = $(patsubst %.txt,%.1,$(MAN1_TXT))
DOC_MAN = $(DOC_MAN1)

ASCIIDOC = asciidoc
ASCIIDOC_EXTRA =
ASCIIDOC_HTML = xhtml11
ASCIIDOC_DOCBOOK = docbook
ASCIIDOC_CONF = -f asciidoc.conf
ASCIIDOC_COMMON = $(ASCIIDOC) $(ASCIIDOC_EXTRA) $(ASCIIDOC_CONF)

TXT_TO_HTML = $(ASCIIDOC_COMMON) -b $(ASCIIDOC_HTML)
TXT_TO_XML = $(ASCIIDOC_COMMON) -b $(ASCIIDOC_DOCBOOK)
XMLTO = xmlto
XMLTO_EXTRA =
INSTALL ?= install
RM ?= rm -f

prefix = /usr
mandir = $(prefix)/share/man
man1dir = $(mandir)/man1

all: html man

html: $(DOC_HTML)

man: $(DOC_MAN)

install:
	$(INSTALL) -d -m 755 $(man1dir)
	$(INSTALL) -m 644 $(DOC_MAN1) $(man1dir)

clean:
	$(RM) *.xml *.html *.1

%.html : %.txt
	$(TXT_TO_HTML) -d manpage -o $@ $<

%.1 : %.xml
	$(XMLTO) $(XMLTO_EXTRA) man $<

%.xml : %.txt asciidoc.conf
	$(TXT_TO_XML) -d manpage -o $@ $<

.PHONY: FORCE
