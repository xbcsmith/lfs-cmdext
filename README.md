# lfs-cmdext

LFS Command Extractor

## Description

a Go LFS Command Extractor used to extract package information from the LFS/BLFS books.

## Build

`make`

## Usage

```
Usage of ./bin/cmdext:
  -asjson
    	Output JSON
  -asyaml
    	Output YAML (default true)
  -debug
    	Turn debugging on
  -destination string
    	Path to write files to disk (default "/tmp/pkgs")
  -write-to-disk
    	Write files to disk

```

## Example

### Fetch BLFS

```bash
curl -O http://www.linuxfromscratch.org/blfs/downloads/9.0-systemd/blfs-book-9.0-systemd-html.tar.xz
```

```bash
tar -xvf blfs-book-9.0-systemd-html.tar.xz
```

### Extract Commands

```
cmdext general/tcl.html
```

### Output

```yaml
commands:
  - cmd: tar -xf ../tcl8.6.9-html.tar.gz --strip-components=1
    index: 0
  - cmd: |-
        export SRCDIR=`pwd` &&

        cd unix &&

        ./configure --prefix=/usr           \
                    --mandir=/usr/share/man \
                    $([ $(uname -m) = x86_64 ] && echo --enable-64bit) &&
        make &&

        sed -e "s#$SRCDIR/unix#/usr/lib#" \
            -e "s#$SRCDIR#/usr/include#"  \
            -i tclConfig.sh               &&

        sed -e "s#$SRCDIR/unix/pkgs/tdbc1.1.0#/usr/lib/tdbc1.1.0#" \
            -e "s#$SRCDIR/pkgs/tdbc1.1.0/generic#/usr/include#"    \
            -e "s#$SRCDIR/pkgs/tdbc1.1.0/library#/usr/lib/tcl8.6#" \
            -e "s#$SRCDIR/pkgs/tdbc1.1.0#/usr/include#"            \
            -i pkgs/tdbc1.1.0/tdbcConfig.sh                        &&

        sed -e "s#$SRCDIR/unix/pkgs/itcl4.1.2#/usr/lib/itcl4.1.2#" \
            -e "s#$SRCDIR/pkgs/itcl4.1.2/generic#/usr/include#"    \
            -e "s#$SRCDIR/pkgs/itcl4.1.2#/usr/include#"            \
            -i pkgs/itcl4.1.2/itclConfig.sh                        &&

        unset SRCDIR
    index: 1
  - cmd: |-
        make install &&
        make install-private-headers &&
        ln -v -sf tclsh8.6 /usr/bin/tclsh &&
        chmod -v 755 /usr/lib/libtcl8.6.so
    index: 2
  - cmd: |-
        mkdir -v -p /usr/share/doc/tcl-8.6.9 &&
        cp -v -r  ../html/* /usr/share/doc/tcl-8.6.9
    index: 3
dependencies:
    optional: []
    recommended: []
    requires: []
description: |-
    The Tcl package contains the Tool
              Command Language, a robust general-purpose scripting language.
name: tcl
sources:
  - archive: https://downloads.sourceforge.net/tcl/tcl8.6.9-src.tar.gz
    build_time: 0.6 SBU (add 3.0 SBU for tests)
    md5sum: aa0a121d95a0e7b73a036f26028538d4
    ondisk: |-
        84 MB (including html
                        documentation)
    size: 9.5 MB
  - archive: https://downloads.sourceforge.net/tcl/tcl8.6.9-html.tar.gz
    build_time: ""
    md5sum: 243da67cca49b9bac0dc6c06fdb42896
    ondisk: ""
    size: 1.2 MB
version: 8.6.9

```

### More

```
cmdext --asjson general/tcl.html

cmdext --asjson --noindent general/tcl.html

cmdext --asjson --write-to-disk general/tcl.html

cmdext --write-to-disk general/tcl.html

cmdext --write-to-disk --destination "./pkgs" general/tcl.html
```
