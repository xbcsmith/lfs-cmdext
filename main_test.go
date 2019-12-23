// Copyright © 2019 Brett Smith <xbcsmith@gmail.com>, . All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

// TestExtractCommands func takes no input and returns t *testing.T
func TestExtractCommands(t *testing.T) {
	htmlPkg := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
    "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta http-equiv="Content-Type" content=
    "application/xhtml+xml; charset=iso-8859-1" />
    <title>
      Exiv2-0.27.2
    </title>
    <link rel="stylesheet" href="../stylesheets/lfs.css" type="text/css" />
    <meta name="generator" content="DocBook XSL Stylesheets V1.73.2" />
    <link rel="stylesheet" href="../stylesheets/lfs-print.css" type=
    "text/css" media="print" />
  </head>
  <body class="blfs" id="blfs-9.0">
    <div class="navheader">
      <h4>
        Beyond Linux<sup>®</sup> From Scratch <span>(systemd</span> Edition)
        - Version 9.0
      </h4>
      <h3>
        Chapter&nbsp;10.&nbsp;Graphics and Font Libraries
      </h3>
      <ul>
        <li class="prev">
          <a accesskey="p" href="babl.html" title="babl-0.1.70">Prev</a>
          <p>
            babl-0.1.70
          </p>
        </li>
        <li class="next">
          <a accesskey="n" href="freetype2.html" title=
          "FreeType-2.10.1">Next</a>
          <p>
            FreeType-2.10.1
          </p>
        </li>
        <li class="up">
          <a accesskey="u" href="graphlib.html" title=
          "Chapter&nbsp;10.&nbsp;Graphics and Font Libraries">Up</a>
        </li>
        <li class="home">
          <a accesskey="h" href="../index.html" title=
          "Beyond Linux® From Scratch   (systemd  Edition) - Version 9.0">Home</a>
        </li>
      </ul>
    </div>
    <div class="sect1" lang="en" xml:lang="en">
      <h1 class="sect1">
        <a id="exiv2" name="exiv2"></a>Exiv2-0.27.2
      </h1>
      <div class="package" lang="en" xml:lang="en">
        <h2 class="sect2">
          Introduction to Exiv2
        </h2>
        <p>
          <span class="application">Exiv2</span> is a C++ library and a
          command line utility for managing image and video metadata.
        </p>
        <p>
          This package is known to build and work properly using an LFS-9.0
          platform.
        </p>
        <h3>
          Package Information
        </h3>
        <div class="itemizedlist">
          <ul class="compact">
            <li>
              <p>
                Download (HTTP): <a class="ulink" href=
                "http://www.exiv2.org/builds/exiv2-0.27.2-Source.tar.gz">http://www.exiv2.org/builds/exiv2-0.27.2-Source.tar.gz</a>
              </p>
            </li>
            <li>
              <p>
                Download MD5 sum: 8c39c39dc8141bb158e8e9d663bcbf21
              </p>
            </li>
            <li>
              <p>
                Download size: 26 MB
              </p>
            </li>
            <li>
              <p>
                Estimated disk space required: 70 MB
              </p>
            </li>
            <li>
              <p>
                Estimated build time: 0.2 SBU (Using parallelism=4)
              </p>
            </li>
          </ul>
        </div>
        <h3>
          Exiv2 dependencies
        </h3>
        <h4>
          Required
        </h4>
        <p class="required">
          <a class="xref" href="cmake.html" title=
          "CMake-3.15.2">CMake-3.15.2</a>
        </p>
        <h4>
          Recommended
        </h4>
        <p class="recommended">
          <a class="xref" href="../basicnet/curl.html" title=
          "cURL-7.65.3">cURL-7.65.3</a>
        </p>
        <h4>
          Optional
        </h4>
        <p class="optional">
          <a class="ulink" href="http://www.libssh.org/">libssh</a>
        </p>
        <h4>
          Optional for documentation
        </h4>
        <p class="optional">
          <a class="xref" href="doxygen.html" title=
          "Doxygen-1.8.16">Doxygen-1.8.16</a>, <a class="xref" href=
          "graphviz.html" title="Graphviz-2.40.1">Graphviz-2.40.1</a>, and
          <a class="xref" href="libxslt.html" title=
          "libxslt-1.1.33">libxslt-1.1.33</a>
        </p>
        <p class="usernotes">
          User Notes: <a class="ulink" href=
          "http://wiki.linuxfromscratch.org/blfs/wiki/exiv2">http://wiki.linuxfromscratch.org/blfs/wiki/exiv2</a>
        </p>
      </div>
      <div class="installation" lang="en" xml:lang="en">
        <h2 class="sect2">
          Installation of Exiv2
        </h2>
        <p>
          Install <span class="application">Exiv2</span> by running the
          following commands:
        </p>
        <pre class="userinput">
<kbd class="command">mkdir build &amp;&amp;
cd    build &amp;&amp;

cmake -DCMAKE_INSTALL_PREFIX=/usr  \
      -DCMAKE_BUILD_TYPE=Release   \
      -DEXIV2_ENABLE_VIDEO=yes     \
      -DEXIV2_ENABLE_WEBREADY=yes  \
      -DEXIV2_ENABLE_CURL=yes      \
      -DEXIV2_BUILD_SAMPLES=no     \
      -G "Unix Makefiles" .. &amp;&amp;
make</kbd>
</pre>
        <p>
          This package does not come with a test suite.
        </p>
        <p>
          Now, as the <code class="systemitem">root</code> user:
        </p>
        <pre class="root">
<kbd class="command">make install</kbd>
</pre>
      </div>
      <div class="commands" lang="en" xml:lang="en">
        <h2 class="sect2">
          Command Explanations
        </h2>
        <p>
          <em class="parameter"><code>-DEXIV2_ENABLE_VIDEO=yes</code></em>:
          This switch enables managing video metadata.
        </p>
        <p>
          <em class=
          "parameter"><code>-DEXIV2_ENABLE_WEBREADY=yes</code></em>: This
          switch enables managing web image metadata.
        </p>
        <p>
          <em class="parameter"><code>-DEXIV2_BUILD_SAMPLES=no</code></em>:
          This switch is necessary to suppress building and installing sample
          programs. If the sample programs are built, 34 additional programs
          are installed in /usr/bin.
        </p>
        <p>
          <em class="parameter"><code>-DEXIV2_ENABLE_CURL=yes</code></em>:
          This switch is necessary to enable network/http capabilities.
        </p>
      </div>
      <div class="content" lang="en" xml:lang="en">
        <h2 class="sect2">
          Contents
        </h2>
        <div class="segmentedlist">
          <div class="seglistitem">
            <div class="seg">
              <strong class="segtitle">Installed Program:</strong>
              <span class="segbody">exiv2</span>
            </div>
            <div class="seg">
              <strong class="segtitle">Installed Library:</strong>
              <span class="segbody">libexiv2.so and libxmp.a</span>
            </div>
            <div class="seg">
              <strong class="segtitle">Installed Directories:</strong>
              <span class="segbody">/usr/include/exiv2 and
              /usr/share/exiv2</span>
            </div>
          </div>
        </div>
        <div class="variablelist">
          <h3>
            Short Descriptions
          </h3>
          <table border="0">
            <col align="left" valign="top" />
            <tbody>
              <tr>
                <td>
                  <p>
                    <a id="exiv2-prog" name="exiv2-prog"></a><span class=
                    "term"><span class=
                    "command"><strong>exiv2</strong></span></span>
                  </p>
                </td>
                <td>
                  <p>
                    is a utility used to dump Exif data.
                  </p>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <p class="updated">
        Last updated on 2019-08-19 11:11:15 -0700
      </p>
    </div>
    <div class="navfooter">
      <ul>
        <li class="prev">
          <a accesskey="p" href="babl.html" title="babl-0.1.70">Prev</a>
          <p>
            babl-0.1.70
          </p>
        </li>
        <li class="next">
          <a accesskey="n" href="freetype2.html" title=
          "FreeType-2.10.1">Next</a>
          <p>
            FreeType-2.10.1
          </p>
        </li>
        <li class="up">
          <a accesskey="u" href="graphlib.html" title=
          "Chapter&nbsp;10.&nbsp;Graphics and Font Libraries">Up</a>
        </li>
        <li class="home">
          <a accesskey="h" href="../index.html" title=
          "Beyond Linux® From Scratch   (systemd  Edition) - Version 9.0">Home</a>
        </li>
      </ul>
    </div>
  </body>
</html>
`
	doc, err := ReadDoc([]byte(htmlPkg))
	assert.Assert(t, is.Nil(err))
	commands, err := ExtractCommands(doc)
	assert.Assert(t, is.Nil(err))
	for _, cmd := range commands {
		fmt.Printf("INDEX : %d\n", cmd.Index)
		fmt.Printf("COMMAND : %s\n", cmd.Cmd)
	}
	sources, err := ExtractSources(doc)
	assert.Assert(t, is.Nil(err))
	for _, src := range sources {
		fmt.Printf("SOURCE ARCHIVE : %s\n", src.Archive)
		fmt.Printf("SOURCE MD5 : %s\n", src.MD5Sum)
		fmt.Printf("SOURCE SIZE : %s\n", src.Size)
		fmt.Printf("SOURCE DISK SPACE : %s\n", src.OnDisk)
		fmt.Printf("SOURCE BUILD TIME : %s\n", src.BuildTime)
	}
	app, err := ExtractApplication(doc)
	assert.Assert(t, is.Nil(err))
	fmt.Printf("APPLICATION NAME : %s\n", app.Name)
	fmt.Printf("APPLICATION VERSION : %s\n", app.Version)
	fmt.Printf("APPLICATION DESCRIPTION : %s\n", app.Description)
	deps, err := ExtractDependencies(doc)
	assert.Assert(t, is.Nil(err))
	for _, req := range deps.Requires {
		fmt.Printf("REQUIRES : %s\n", req)
	}
	for _, rec := range deps.Recommended {
		fmt.Printf("RECOMMENDED : %s\n", rec)
	}
	for _, opt := range deps.Optional {
		fmt.Printf("OPTIONAL : %s\n", opt)
	}
}

// TestExtractApplication func takes no input and returns t *testing.T
func TestExtractSources(t *testing.T) {
	htmlPkg := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
    "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta http-equiv="Content-Type" content=
    "application/xhtml+xml; charset=iso-8859-1" />
    <title>
      FreeType-2.10.1
    </title>
    <link rel="stylesheet" href="../stylesheets/lfs.css" type="text/css" />
    <meta name="generator" content="DocBook XSL Stylesheets V1.73.2" />
    <link rel="stylesheet" href="../stylesheets/lfs-print.css" type=
    "text/css" media="print" />
  </head>
  <body class="blfs" id="blfs-9.0">
    <div class="navheader">
      <h4>
        Beyond Linux<sup>®</sup> From Scratch <span>(systemd</span> Edition)
        - Version 9.0
      </h4>
      <h3>
        Chapter&nbsp;10.&nbsp;Graphics and Font Libraries
      </h3>
      <ul>
        <li class="prev">
          <a accesskey="p" href="exiv2.html" title="Exiv2-0.27.2">Prev</a>
          <p>
            Exiv2-0.27.2
          </p>
        </li>
        <li class="next">
          <a accesskey="n" href="fontconfig.html" title=
          "Fontconfig-2.13.1">Next</a>
          <p>
            Fontconfig-2.13.1
          </p>
        </li>
        <li class="up">
          <a accesskey="u" href="graphlib.html" title=
          "Chapter&nbsp;10.&nbsp;Graphics and Font Libraries">Up</a>
        </li>
        <li class="home">
          <a accesskey="h" href="../index.html" title=
          "Beyond Linux® From Scratch   (systemd  Edition) - Version 9.0">Home</a>
        </li>
      </ul>
    </div>
    <div class="sect1" lang="en" xml:lang="en">
      <h1 class="sect1">
        <a id="freetype2" name="freetype2"></a>FreeType-2.10.1
      </h1>
      <div class="package" lang="en" xml:lang="en">
        <h2 class="sect2">
          Introduction to FreeType2
        </h2>
        <p>
          The <span class="application">FreeType2</span> package contains a
          library which allows applications to properly render <span class=
          "application">TrueType</span> fonts.
        </p>
        <p>
          This package is known to build and work properly using an LFS-9.0
          platform.
        </p>
        <h3>
          Package Information
        </h3>
        <div class="itemizedlist">
          <ul class="compact">
            <li>
              <p>
                Download (HTTP): <a class="ulink" href=
                "https://downloads.sourceforge.net/freetype/freetype-2.10.1.tar.xz">
                https://downloads.sourceforge.net/freetype/freetype-2.10.1.tar.xz</a>
              </p>
            </li>
            <li>
              <p>
                Download MD5 sum: bd42e75127f8431923679480efb5ba8f
              </p>
            </li>
            <li>
              <p>
                Download size: 2.3 MB
              </p>
            </li>
            <li>
              <p>
                Estimated disk space required: 30 MB (with additional
                documentation)
              </p>
            </li>
            <li>
              <p>
                Estimated build time: 0.2 SBU (with additional documentation)
              </p>
            </li>
          </ul>
        </div>
        <h3>
          Additional Downloads
        </h3>
        <div class="itemizedlist">
          <p class="title">
            <b>Additional Documentation</b>
          </p>
          <ul class="compact">
            <li>
              <p>
                Download (HTTP): <a class="ulink" href=
                "https://downloads.sourceforge.net/freetype/freetype-doc-2.10.1.tar.xz">
                https://downloads.sourceforge.net/freetype/freetype-doc-2.10.1.tar.xz</a>
              </p>
            </li>
            <li>
              <p>
                Download MD5 sum: b56d1af90510f0ae4bf12a82410985f5
              </p>
            </li>
            <li>
              <p>
                Download size: 2.0 MB
              </p>
            </li>
          </ul>
        </div>
        <h3>
          FreeType2 Dependencies
        </h3>
        <h4>
          Recommended
        </h4>
        <p class="recommended">
          <a class="xref" href="harfbuzz.html" title=
          "HarfBuzz-2.6.0">HarfBuzz-2.6.0</a> (first, install without it,
          after it is installed, reinstall <a class="xref" href=
          "freetype2.html" title="FreeType-2.10.1">FreeType-2.10.1</a>),
          <a class="xref" href="libpng.html" title=
          "libpng-1.6.37">libpng-1.6.37</a>, and <a class="xref" href=
          "which.html" title="Which-2.21 and Alternatives">Which-2.21</a>
        </p>
        <p class="usernotes">
          User Notes: <a class="ulink" href=
          "http://wiki.linuxfromscratch.org/blfs/wiki/freetype2">http://wiki.linuxfromscratch.org/blfs/wiki/freetype2</a>
        </p>
      </div>
      <div class="installation" lang="en" xml:lang="en">
        <h2 class="sect2">
          Installation of FreeType2
        </h2>
        <p>
          If you downloaded the additional documentation, unpack it into the
          source tree using the following command:
        </p>
        <pre class="userinput">
<kbd class=
"command">tar -xf ../freetype-doc-2.10.1.tar.xz --strip-components=2 -C docs</kbd>
</pre>
        <p>
          Install <span class="application">FreeType2</span> by running the
          following commands:
        </p>
        <pre class="userinput">
<kbd class=
"command">sed -ri "s:.*(AUX_MODULES.*valid):\1:" modules.cfg &amp;&amp;

sed -r "s:.*(#.*SUBPIXEL_RENDERING) .*:\1:" \
    -i include/freetype/config/ftoption.h  &amp;&amp;

./configure --prefix=/usr --enable-freetype-config --disable-static &amp;&amp;
make</kbd>
</pre>
        <p>
          This package does not come with a test suite.
        </p>
        <p>
          Now, as the <code class="systemitem">root</code> user:
        </p>
        <pre class="root">
<kbd class="command">make install</kbd>
</pre>
        <p>
          If you downloaded the optional documentation, install it as the
          <code class="systemitem">root</code> user:
        </p>
        <pre class="root">
<kbd class=
"command">install -v -m755 -d /usr/share/doc/freetype-2.10.1 &amp;&amp;
cp -v -R docs/*     /usr/share/doc/freetype-2.10.1 &amp;&amp;
rm -v /usr/share/doc/freetype-2.10.1/freetype-config.1</kbd>
</pre>
      </div>
      <div class="commands" lang="en" xml:lang="en">
        <h2 class="sect2">
          Command Explanations
        </h2>
        <p>
          <span class="command"><strong>sed -ri ...</strong></span>: First
          command enables GX/AAT and OpenType table validation and second
          command enables Subpixel Rendering. Note that Subpixel Rendering
          may have patent issues. Be sure to read the <code class=
          "literal">'Other patent issues'</code> part of <a class="ulink"
          href=
          "http://www.freetype.org/patents.html">http://www.freetype.org/patents.html</a>
          before enabling this option.
        </p>
        <p>
          <em class="parameter"><code>--enable-freetype-config</code></em>:
          This switch ensure that the man page for freetype-config is
          installed.
        </p>
        <p>
          <em class="parameter"><code>--without-harfbuzz</code></em>: If
          <span class="application">harfbuzz</span> is installed prior to
          <span class="application">freetype</span> without <span class=
          "application">freetype</span> support, use this switch to avoid a
          build failure.
        </p>
        <p>
          <em class="parameter"><code>--disable-static</code></em>: This
          switch prevents installation of static versions of the libraries.
        </p>
        <p>
          <span class="command"><strong>cp builds/unix/freetype-config
          /usr/bin</strong></span>: Manually place the freetype
          configureation program needed by other programs when using the
          freetype library.
        </p>
      </div>
      <div class="content" lang="en" xml:lang="en">
        <h2 class="sect2">
          Contents
        </h2>
        <div class="segmentedlist">
          <div class="seglistitem">
            <div class="seg">
              <strong class="segtitle">Installed Program:</strong>
              <span class="segbody">freetype-config</span>
            </div>
            <div class="seg">
              <strong class="segtitle">Installed Library:</strong>
              <span class="segbody">libfreetype.so</span>
            </div>
            <div class="seg">
              <strong class="segtitle">Installed Directories:</strong>
              <span class="segbody">/usr/include/freetype2 and
              /usr/share/doc/freetype-2.10.1</span>
            </div>
          </div>
        </div>
        <div class="variablelist">
          <h3>
            Short Descriptions
          </h3>
          <table border="0">
            <col align="left" valign="top" />
            <tbody>
              <tr>
                <td>
                  <p>
                    <a id="freetype-config" name=
                    "freetype-config"></a><span class="term"><span class=
                    "command"><strong>freetype-config</strong></span></span>
                  </p>
                </td>
                <td>
                  <p>
                    is used to get <span class="application">FreeType</span>
                    compilation and linking information.
                  </p>
                </td>
              </tr>
              <tr>
                <td>
                  <p>
                    <a id="libfreetype" name="libfreetype"></a><span class=
                    "term"><code class=
                    "filename">libfreetype.so</code></span>
                  </p>
                </td>
                <td>
                  <p>
                    contains functions for rendering various font types, such
                    as TrueType and Type1.
                  </p>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <p class="updated">
        Last updated on 2019-08-16 15:28:01 -0700
      </p>
    </div>
    <div class="navfooter">
      <ul>
        <li class="prev">
          <a accesskey="p" href="exiv2.html" title="Exiv2-0.27.2">Prev</a>
          <p>
            Exiv2-0.27.2
          </p>
        </li>
        <li class="next">
          <a accesskey="n" href="fontconfig.html" title=
          "Fontconfig-2.13.1">Next</a>
          <p>
            Fontconfig-2.13.1
          </p>
        </li>
        <li class="up">
          <a accesskey="u" href="graphlib.html" title=
          "Chapter&nbsp;10.&nbsp;Graphics and Font Libraries">Up</a>
        </li>
        <li class="home">
          <a accesskey="h" href="../index.html" title=
          "Beyond Linux® From Scratch   (systemd  Edition) - Version 9.0">Home</a>
        </li>
      </ul>
    </div>
  </body>
</html>
`
	doc, err := ReadDoc([]byte(htmlPkg))
	assert.Assert(t, is.Nil(err))
	commands, err := ExtractCommands(doc)
	assert.Assert(t, is.Nil(err))
	for _, cmd := range commands {
		fmt.Printf("INDEX : %d\n", cmd.Index)
		fmt.Printf("COMMAND : %s\n", cmd.Cmd)
	}
	sources, err := ExtractSources(doc)
	assert.Assert(t, is.Nil(err))
	for _, src := range sources {
		fmt.Printf("SOURCE ARCHIVE : %s\n", src.Archive)
		fmt.Printf("SOURCE MD5 : %s\n", src.MD5Sum)
		fmt.Printf("SOURCE SIZE : %s\n", src.Size)
		fmt.Printf("SOURCE DISK SPACE : %s\n", src.OnDisk)
		fmt.Printf("SOURCE BUILD TIME : %s\n", src.BuildTime)
	}
	app, err := ExtractApplication(doc)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, app.Name, "freetype")
	assert.Equal(t, app.Version, "2.10.1")
	fmt.Printf("APPLICATION NAME : %s\n", app.Name)
	fmt.Printf("APPLICATION VERSION : %s\n", app.Version)
	fmt.Printf("APPLICATION DESCRIPTION : %s\n", app.Description)
	deps, err := ExtractDependencies(doc)
	assert.Assert(t, is.Nil(err))
	for _, req := range deps.Requires {
		fmt.Printf("REQUIRES : %s\n", req)
	}
	for _, rec := range deps.Recommended {
		fmt.Printf("RECOMMENDED : %s\n", rec)
	}
	for _, opt := range deps.Optional {
		fmt.Printf("OPTIONAL : %s\n", opt)
	}
	pkg, err := CreatePackageInformation([]byte(htmlPkg))
	assert.Assert(t, is.Nil(err))
	yml, err := pkg.ToYAML()
	assert.Assert(t, is.Nil(err))
	fmt.Printf("%s\n", yml)
}

// TestExtractCommands func takes no input and returns t *testing.T
func TestExtractApplication(t *testing.T) {
	htmlPkg := `<?xml version="1.0" encoding="iso-8859-1" standalone="no"?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
    "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta http-equiv="Content-Type" content=
    "application/xhtml+xml; charset=iso-8859-1" />
    <title>
      Configuring the JAVA environment
    </title>
    <link rel="stylesheet" type="text/css" href="../stylesheets/lfs.css" />
    <meta name="generator" content="DocBook XSL Stylesheets V1.78.1" />
    <link rel="stylesheet" href="../stylesheets/lfs-print.css" type=
    "text/css" media="print" />
  </head>
  <body class="blfs" id="blfs-2019-12-22">
    <div class="navheader">
      <h4>
        Beyond Linux<sup>®</sup> From Scratch <span class="phrase">(System
        V</span> Edition) - Version 2019-12-22
      </h4>
      <h3>
        Chapter&nbsp;13.&nbsp;Programming
      </h3>
      <ul>
        <li class="prev">
          <a accesskey="p" href="openjdk.html" title=
          "OpenJDK-12.0.2">Prev</a>
          <p>
            OpenJDK-12.0.2
          </p>
        </li>
        <li class="next">
          <a accesskey="n" href="apache-ant.html" title=
          "apache-ant-1.10.7">Next</a>
          <p>
            apache-ant-1.10.7
          </p>
        </li>
        <li class="up">
          <a accesskey="u" href="prog.html" title=
          "Chapter&nbsp;13.&nbsp;Programming">Up</a>
        </li>
        <li class="home">
          <a accesskey="h" href="../index.html" title=
          "Beyond Linux® From Scratch     (System V Edition) - Version 2019-12-22">
          Home</a>
        </li>
      </ul>
    </div>
    <div class="sect1" lang="en" xml:lang="en">
      <h1 class="sect1">
        <a id="ojdk-conf" name="ojdk-conf"></a>Configuring the JAVA
        environment
      </h1>
      <div class="sect2" lang="en" xml:lang="en">
        <h2 class="sect2">
          <a id="java-profile" name="java-profile"></a>Setting up the
          environment
        </h2>
        <p>
          After the package installation is complete, the next step is to
          make sure that the system can properly find the files. If you set
          up your login scripts as recommended in <a class="xref" href=
          "../postlfs/profile.html" title="The Bash Shell Startup Files">The
          Bash Shell Startup Files</a>, update the environment by creating
          the <code class="filename">openjdk.sh</code> script, as the
          <code class="systemitem">root</code> user:
        </p>
        <pre class="root">
<kbd class="command">cat &gt; /etc/profile.d/openjdk.sh &lt;&lt; "EOF"
<code class="literal"># Begin /etc/profile.d/openjdk.sh

# Set JAVA_HOME directory
JAVA_HOME=/opt/jdk

# Adjust PATH
pathappend $JAVA_HOME/bin

# Add to MANPATH
pathappend $JAVA_HOME/man MANPATH

# Auto Java CLASSPATH: Copy jar files to, or create symlinks in, the
# /usr/share/java directory. Note that having gcj jars with OpenJDK 8
# may lead to errors.

AUTO_CLASSPATH_DIR=/usr/share/java

pathprepend . CLASSPATH

for dir in $(find ${AUTO_CLASSPATH_DIR} -type d 2&gt;/dev/null); do
    pathappend $dir CLASSPATH
done

for jar in $(find ${AUTO_CLASSPATH_DIR} -name "*.jar" 2&gt;/dev/null); do
    pathappend $jar CLASSPATH
done

export JAVA_HOME
unset AUTO_CLASSPATH_DIR dir jar

# End /etc/profile.d/openjdk.sh</code>
EOF</kbd>
</pre>
        <p>
          If <a class="xref" href="../postlfs/sudo.html" title=
          "Sudo-1.8.29">Sudo-1.8.29</a> is installed, the super user should
          have access to the above variables. Execute the following commands
          as the <code class="systemitem">root</code> user:
        </p>
        <pre class="root">
<kbd class="command">cat &gt; /etc/sudoers.d/java &lt;&lt; "EOF"
<code class="literal">Defaults env_keep += JAVA_HOME
Defaults env_keep += CLASSPATH</code>
EOF</kbd>
</pre>
        <p>
          For allowing <span class="command"><strong>mandb</strong></span> to
          include the OpenJDK man pages in its database, issue, as the
          <code class="systemitem">root</code> user:
        </p>
        <pre class="root">
<kbd class="command">cat &gt;&gt; /etc/man_db.conf &lt;&lt; "EOF" &amp;&amp;
<code class="literal"># Begin Java addition
MANDATORY_MANPATH     /opt/jdk/man
MANPATH_MAP           /opt/jdk/bin     /opt/jdk/man
MANDB_MAP             /opt/jdk/man     /var/cache/man/jdk
# End Java addition</code>
EOF

mkdir -p /var/cache/man &amp;&amp;
mandb -c /opt/jdk/man</kbd>
</pre>
      </div>
      <div class="sect2" lang="en" xml:lang="en">
        <h2 class="sect2">
          <a id="ojdk-certs" name="ojdk-certs"></a>Setting up the Certificate
          Authority Certificates for Java
        </h2>
        <p>
          <span class="application">OpenJDK</span> uses its own format for
          the CA certificates. The Java security modules use <code class=
          "envar">$JAVA_HOME</code><code class=
          "filename">/lib/security/cacerts</code> by default. In order to
          keep all the certificates in one place, we use <code class=
          "filename">/etc/ssl/java/cacerts</code>. The instructions on the
          <a class="xref" href="../postlfs/make-ca.html" title=
          "make-ca-1.5">make-ca-1.5</a> page previously created the file
          located in <code class="filename">/etc/ssl/java</code>. Setup a
          symlink in the default location as the <code class=
          "systemitem">root</code> user:
        </p>
        <pre class="root">
<kbd class=
"command">ln -sfv /etc/pki/tls/java/cacerts /opt/jdk/lib/security/cacerts</kbd>
</pre>
        <p>
          Use the following command to check if the <code class=
          "filename">cacerts</code> file has been successfully installed:
        </p>
        <pre class="root">
<kbd class="command">/opt/jdk/bin/keytool -list -cacerts</kbd>
</pre>
        <p>
          At the prompt <code class="computeroutput">Enter keystore
          password:</code>, enter <strong class=
          "userinput"><code>changeit</code></strong> (the default) or just
          press the <span class="quote">&ldquo;<span class=
          "quote">Enter</span>&rdquo;</span> key. If the <code class=
          "filename">cacerts</code> file was installed correctly, you will
          see a list of the certificates with related information for each
          one. If not, you need to reinstall them.
        </p>
        <p>
          If you later install a new JVM, you just have to create the symlink
          in the default location to be able to use the cacerts.
        </p>
      </div>
      <p class="updated">
        Last updated on 2019-02-09 20:26:31 -0600
      </p>
    </div>
    <div class="navfooter">
      <ul>
        <li class="prev">
          <a accesskey="p" href="openjdk.html" title=
          "OpenJDK-12.0.2">Prev</a>
          <p>
            OpenJDK-12.0.2
          </p>
        </li>
        <li class="next">
          <a accesskey="n" href="apache-ant.html" title=
          "apache-ant-1.10.7">Next</a>
          <p>
            apache-ant-1.10.7
          </p>
        </li>
        <li class="up">
          <a accesskey="u" href="prog.html" title=
          "Chapter&nbsp;13.&nbsp;Programming">Up</a>
        </li>
        <li class="home">
          <a accesskey="h" href="../index.html" title=
          "Beyond Linux® From Scratch     (System V Edition) - Version 2019-12-22">
          Home</a>
        </li>
      </ul>
    </div>
  </body>
</html>
`

	doc, err := ReadDoc([]byte(htmlPkg))
	assert.Assert(t, is.Nil(err))
	commands, err := ExtractCommands(doc)
	assert.Assert(t, is.Nil(err))
	for _, cmd := range commands {
		fmt.Printf("INDEX : %d\n", cmd.Index)
		fmt.Printf("COMMAND : %s\n", cmd.Cmd)
	}
	sources, err := ExtractSources(doc)
	assert.Assert(t, is.Nil(err))
	for _, src := range sources {
		fmt.Printf("SOURCE ARCHIVE : %s\n", src.Archive)
		fmt.Printf("SOURCE MD5 : %s\n", src.MD5Sum)
		fmt.Printf("SOURCE SIZE : %s\n", src.Size)
		fmt.Printf("SOURCE DISK SPACE : %s\n", src.OnDisk)
		fmt.Printf("SOURCE BUILD TIME : %s\n", src.BuildTime)
	}
	app, err := ExtractApplication(doc)
	assert.Assert(t, is.Nil(err))
	fmt.Printf("APPLICATION NAME : %s\n", app.Name)
	fmt.Printf("APPLICATION VERSION : %s\n", app.Version)
	fmt.Printf("APPLICATION DESCRIPTION : %s\n", app.Description)
	deps, err := ExtractDependencies(doc)
	assert.Assert(t, is.Nil(err))
	for _, req := range deps.Requires {
		fmt.Printf("REQUIRES : %s\n", req)
	}
	for _, rec := range deps.Recommended {
		fmt.Printf("RECOMMENDED : %s\n", rec)
	}
	for _, opt := range deps.Optional {
		fmt.Printf("OPTIONAL : %s\n", opt)
	}
	pkg, err := CreatePackageInformation([]byte(htmlPkg))
	assert.Assert(t, is.Nil(err))
	yml, err := pkg.ToYAML()
	assert.Assert(t, is.Nil(err))
	fmt.Printf("%s\n", yml)
}
