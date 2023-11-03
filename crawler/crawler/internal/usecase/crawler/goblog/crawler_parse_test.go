package goblog

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

var body = `
<!DOCTYPE html>
<html lang="en" data-theme="auto">
<head>

<link rel="preconnect" href="https://www.googletagmanager.com">
<script >(function(w,d,s,l,i){w[l]=w[l]||[];w[l].push({'gtm.start':
  new Date().getTime(),event:'gtm.js'});var f=d.getElementsByTagName(s)[0],
  j=d.createElement(s),dl=l!='dataLayer'?'&l='+l:'';j.async=true;j.src=
  'https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);
  })(window,document,'script','dataLayer','GTM-W8MVQXG');</script>
  
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="theme-color" content="#00add8">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Material+Icons">
<link rel="stylesheet" href="/css/styles.css">
<link rel="icon" href="/images/favicon-gopher.png" sizes="any">
<link rel="apple-touch-icon" href="/images/favicon-gopher-plain.png"/>
<link rel="icon" href="/images/favicon-gopher.svg" type="image/svg+xml">

<link rel="alternate" title="The Go Blog" type="application/atom+xml" href="/blog/feed.atom">

  
  <script>(function(w,d,s,l,i){w[l]=w[l]||[];w[l].push({'gtm.start':
  new Date().getTime(),event:'gtm.js'});var f=d.getElementsByTagName(s)[0],
  j=d.createElement(s),dl=l!='dataLayer'?'&l='+l:'';j.async=true;j.src=
  'https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);
  })(window,document,'script','dataLayer','GTM-W8MVQXG');</script>
  
<script src="/js/site.js"></script>
<meta name="og:url" content="https://go.dev/blog/">
<meta name="og:title" content="The Go Blog - The Go Programming Language">
<title>The Go Blog - The Go Programming Language</title>

<meta name="og:image" content="https://go.dev/doc/gopher/gopher5logo.jpg">
<meta name="twitter:image" content="https://go.dev/doc/gopher/runningsquare.jpg">
<meta name="twitter:card" content="summary">
<meta name="twitter:site" content="@golang">
</head>
<body class="Site">
  
<noscript><iframe src="https://www.googletagmanager.com/ns.html?id=GTM-W8MVQXG"
  height="0" width="0" style="display:none;visibility:hidden"></iframe></noscript>
  


<header class="Site-header js-siteHeader">
  <div class="Header Header--dark">
    <nav class="Header-nav">
      <a href="/">
        <img
          class="js-headerLogo Header-logo"
          src="/images/go-logo-white.svg"
          alt="Go">
      </a>
      <div class="skip-navigation-wrapper">
        <a class="skip-to-content-link" aria-label="Skip to main content" href="#main-content"> Skip to Main Content </a>
      </div>
      <div class="Header-rightContent">
        <ul class="Header-menu">
          <li class="Header-menuItem ">
            <a href="#"  class="js-desktop-menu-hover" aria-label=Why&#32;Go aria-describedby="dropdown-description">
              Why Go <i class="material-icons" aria-hidden="true">arrow_drop_down</i>
            </a>
            <div class="screen-reader-only" id="dropdown-description" hidden> 
              Press Enter to activate/deactivate dropdown
            </div>
              <ul class="Header-submenu js-desktop-submenu-hover" aria-label="submenu">
                  <li class="Header-submenuItem">
                    <div>
                        <a href="/solutions/case-studies">
                          Case Studies
                          
                        </a>
                    </div>
                    <p>Common problems companies solve with Go</p>
                  </li>
                  <li class="Header-submenuItem">
                    <div>
                        <a href="/solutions/use-cases">
                          Use Cases
                          
                        </a>
                    </div>
                    <p>Stories about how and why companies use Go</p>
                  </li>
                  <li class="Header-submenuItem">
                    <div>
                        <a href="/security/">
                          Security
                          
                        </a>
                    </div>
                    <p>How Go can help keep you secure by default</p>
                  </li>
              </ul>
          </li>
          <li class="Header-menuItem ">
            <a href="/learn/"  aria-label=Learn aria-describedby="dropdown-description">
              Learn 
            </a>
            <div class="screen-reader-only" id="dropdown-description" hidden> 
              Press Enter to activate/deactivate dropdown
            </div>
          </li>
          <li class="Header-menuItem ">
            <a href="#"  class="js-desktop-menu-hover" aria-label=Docs aria-describedby="dropdown-description">
              Docs <i class="material-icons" aria-hidden="true">arrow_drop_down</i>
            </a>
            <div class="screen-reader-only" id="dropdown-description" hidden> 
              Press Enter to activate/deactivate dropdown
            </div>
              <ul class="Header-submenu js-desktop-submenu-hover" aria-label="submenu">
                  <li class="Header-submenuItem">
                    <div>
                        <a href="/doc/effective_go">
                          Effective Go
                          
                        </a>
                    </div>
                    <p>Tips for writing clear, performant, and idiomatic Go code</p>
                  </li>
                  <li class="Header-submenuItem">
                    <div>
                        <a href="/doc">
                          Go User Manual
                          
                        </a>
                    </div>
                    <p>A complete introduction to building software with Go</p>
                  </li>
                  <li class="Header-submenuItem">
                    <div>
                        <a href="https://pkg.go.dev/std">
                          Standard library
                          
                        </a>
                    </div>
                    <p>Reference documentation for Go&#39;s standard library</p>
                  </li>
                  <li class="Header-submenuItem">
                    <div>
                        <a href="/doc/devel/release">
                          Release Notes
                          
                        </a>
                    </div>
                    <p>Learn what&#39;s new in each Go release</p>
                  </li>
              </ul>
          </li>
          <li class="Header-menuItem ">
            <a href="https://pkg.go.dev"  aria-label=Packages aria-describedby="dropdown-description">
              Packages 
            </a>
            <div class="screen-reader-only" id="dropdown-description" hidden> 
              Press Enter to activate/deactivate dropdown
            </div>
          </li>
          <li class="Header-menuItem ">
            <a href="#"  class="js-desktop-menu-hover" aria-label=Community aria-describedby="dropdown-description">
              Community <i class="material-icons" aria-hidden="true">arrow_drop_down</i>
            </a>
            <div class="screen-reader-only" id="dropdown-description" hidden> 
              Press Enter to activate/deactivate dropdown
            </div>
              <ul class="Header-submenu js-desktop-submenu-hover" aria-label="submenu">
                  <li class="Header-submenuItem">
                    <div>
                        <a href="/talks/">
                          Recorded Talks
                          
                        </a>
                    </div>
                    <p>Videos from prior events</p>
                  </li>
                  <li class="Header-submenuItem">
                    <div>
                        <a href="https://www.meetup.com/pro/go">
                          Meetups
                           <i class="material-icons">open_in_new</i>
                        </a>
                    </div>
                    <p>Meet other local Go developers</p>
                  </li>
                  <li class="Header-submenuItem">
                    <div>
                        <a href="https://github.com/golang/go/wiki/Conferences">
                          Conferences
                           <i class="material-icons">open_in_new</i>
                        </a>
                    </div>
                    <p>Learn and network with Go developers from around the world</p>
                  </li>
                  <li class="Header-submenuItem">
                    <div>
                        <a href="/blog">
                          Go blog
                          
                        </a>
                    </div>
                    <p>The Go project&#39;s official blog.</p>
                  </li>
                  <li class="Header-submenuItem">
                    <div>
                        <a href="/help">
                          Go project
                          
                        </a>
                    </div>
                    <p>Get help and stay informed from Go</p>
                  </li>
                  <li class="Header-submenuItem">
                    <div>
                        Get connected
                    </div>
                    <p></p>
                      <div class="Header-socialIcons">
                        
                        <a class="Header-socialIcon" aria-label="Get connected with google-groups (Opens in new window)" href="https://groups.google.com/g/golang-nuts"><img src="/images/logos/social/google-groups.svg" /></a>
                        <a class="Header-socialIcon" aria-label="Get connected with github (Opens in new window)" href="https://github.com/golang"><img src="/images/logos/social/github.svg" /></a>
                        <a class="Header-socialIcon" aria-label="Get connected with twitter (Opens in new window)" href="https://twitter.com/golang"><img src="/images/logos/social/twitter.svg" /></a>
                        <a class="Header-socialIcon" aria-label="Get connected with reddit (Opens in new window)" href="https://www.reddit.com/r/golang/"><img src="/images/logos/social/reddit.svg" /></a>
                        <a class="Header-socialIcon" aria-label="Get connected with slack (Opens in new window)" href="https://invite.slack.golangbridge.org/"><img src="/images/logos/social/slack.svg" /></a>
                        <a class="Header-socialIcon" aria-label="Get connected with stack-overflow (Opens in new window)" href="https://stackoverflow.com/tags/go"><img src="/images/logos/social/stack-overflow.svg" /></a>
                      </div>
                  </li>
              </ul>
          </li>
        </ul>
        <button class="Header-navOpen js-headerMenuButton Header-navOpen--white" aria-label="Open navigation.">
        </button>
      </div>
    </nav>
    
  </div>
</header>
<aside class="NavigationDrawer js-header">
  <nav class="NavigationDrawer-nav">
    <div class="NavigationDrawer-header">
      <a href="/">
        <img class="NavigationDrawer-logo" src="/images/go-logo-blue.svg" alt="Go.">
      </a>
    </div>
    <ul class="NavigationDrawer-list">
        
          <li class="NavigationDrawer-listItem js-mobile-subnav-trigger  NavigationDrawer-hasSubnav">
            <a href="#"><span>Why Go</span> <i class="material-icons">navigate_next</i></a>

            <div class="NavigationDrawer NavigationDrawer-submenuItem">
              <nav class="NavigationDrawer-nav">
                <div class="NavigationDrawer-header">
                  <a href="#"><i class="material-icons">navigate_before</i>Why Go</a>
                </div>
                <ul class="NavigationDrawer-list">
                    <li class="NavigationDrawer-listItem">
                        <a href="/solutions/case-studies">
                          Case Studies
                          
                        </a>
                      
                    </li>
                    <li class="NavigationDrawer-listItem">
                        <a href="/solutions/use-cases">
                          Use Cases
                          
                        </a>
                      
                    </li>
                    <li class="NavigationDrawer-listItem">
                        <a href="/security/">
                          Security
                          
                        </a>
                      
                    </li>
                </ul>
              </div>
            </div>
          </li>

        
        
          <li class="NavigationDrawer-listItem ">
            <a href="/learn/">Learn</a>
          </li>
        
        
          <li class="NavigationDrawer-listItem js-mobile-subnav-trigger  NavigationDrawer-hasSubnav">
            <a href="#"><span>Docs</span> <i class="material-icons">navigate_next</i></a>

            <div class="NavigationDrawer NavigationDrawer-submenuItem">
              <nav class="NavigationDrawer-nav">
                <div class="NavigationDrawer-header">
                  <a href="#"><i class="material-icons">navigate_before</i>Docs</a>
                </div>
                <ul class="NavigationDrawer-list">
                    <li class="NavigationDrawer-listItem">
                        <a href="/doc/effective_go">
                          Effective Go
                          
                        </a>
                      
                    </li>
                    <li class="NavigationDrawer-listItem">
                        <a href="/doc">
                          Go User Manual
                          
                        </a>
                      
                    </li>
                    <li class="NavigationDrawer-listItem">
                        <a href="https://pkg.go.dev/std">
                          Standard library
                          
                        </a>
                      
                    </li>
                    <li class="NavigationDrawer-listItem">
                        <a href="/doc/devel/release">
                          Release Notes
                          
                        </a>
                      
                    </li>
                </ul>
              </div>
            </div>
          </li>

        
        
          <li class="NavigationDrawer-listItem ">
            <a href="https://pkg.go.dev">Packages</a>
          </li>
        
        
          <li class="NavigationDrawer-listItem js-mobile-subnav-trigger  NavigationDrawer-hasSubnav">
            <a href="#"><span>Community</span> <i class="material-icons">navigate_next</i></a>

            <div class="NavigationDrawer NavigationDrawer-submenuItem">
              <nav class="NavigationDrawer-nav">
                <div class="NavigationDrawer-header">
                  <a href="#"><i class="material-icons">navigate_before</i>Community</a>
                </div>
                <ul class="NavigationDrawer-list">
                    <li class="NavigationDrawer-listItem">
                        <a href="/talks/">
                          Recorded Talks
                          
                        </a>
                      
                    </li>
                    <li class="NavigationDrawer-listItem">
                        <a href="https://www.meetup.com/pro/go">
                          Meetups
                           <i class="material-icons">open_in_new</i>
                        </a>
                      
                    </li>
                    <li class="NavigationDrawer-listItem">
                        <a href="https://github.com/golang/go/wiki/Conferences">
                          Conferences
                           <i class="material-icons">open_in_new</i>
                        </a>
                      
                    </li>
                    <li class="NavigationDrawer-listItem">
                        <a href="/blog">
                          Go blog
                          
                        </a>
                      
                    </li>
                    <li class="NavigationDrawer-listItem">
                        <a href="/help">
                          Go project
                          
                        </a>
                      
                    </li>
                    <li class="NavigationDrawer-listItem">
                        <div>Get connected</div>
                        <div class="Header-socialIcons">
                          
                            <a class="Header-socialIcon" href="https://groups.google.com/g/golang-nuts"><img src="/images/logos/social/google-groups.svg" /></a>
                            <a class="Header-socialIcon" href="https://github.com/golang"><img src="/images/logos/social/github.svg" /></a>
                            <a class="Header-socialIcon" href="https://twitter.com/golang"><img src="/images/logos/social/twitter.svg" /></a>
                            <a class="Header-socialIcon" href="https://www.reddit.com/r/golang/"><img src="/images/logos/social/reddit.svg" /></a>
                            <a class="Header-socialIcon" href="https://invite.slack.golangbridge.org/"><img src="/images/logos/social/slack.svg" /></a>
                            <a class="Header-socialIcon" href="https://stackoverflow.com/tags/go"><img src="/images/logos/social/stack-overflow.svg" /></a>
                        </div>
                    </li>
                </ul>
              </div>
            </div>
          </li>

        
    </ul>
  </nav>
</aside>
<div class="NavigationDrawer-scrim js-scrim" role="presentation"></div>
<main class="SiteContent SiteContent--default" id="main-content">
  
<div id="blog"><div id="content">
  <div id="content">

    <div class="Article" data-slug="/blog/">
    

    <h1>The Go Blog</h1>
      
      <div id="blogindex">
<p class="blogtitle">
  <a href="/blog/type-inference" aria-describedby="blog-description">Everything You Always Wanted to Know About Type Inference - And a Little Bit More</a>, <span class="date">9 October 2023</span><br>
  <span class="author">Robert Griesemer<br></span>
</p>
<p class="blogsummary">
  A description of how type inference for Go works. Based on the GopherCon 2023 talk with the same title.
</p>
<p class="blogtitle">
  <a href="/blog/deconstructing-type-parameters" aria-describedby="blog-description">Deconstructing Type Parameters</a>, <span class="date">26 September 2023</span><br>
  <span class="author">Ian Lance Taylor<br></span>
</p>
<p class="blogsummary">
  Why the function signatures in the slices packages are so complicated.
</p>
<p class="blogtitle">
  <a href="/blog/loopvar-preview" aria-describedby="blog-description">Fixing For Loops in Go 1.22</a>, <span class="date">19 September 2023</span><br>
  <span class="author">David Chase and Russ Cox<br></span>
</p>
<p class="blogsummary">
  Go 1.21 shipped a preview of a change in Go 1.22 to make for loops less error-prone.
</p>
<p class="blogtitle">
  <a href="/blog/wasi" aria-describedby="blog-description">WASI support in Go</a>, <span class="date">13 September 2023</span><br>
  <span class="author">Johan Brandhorst-Satzkorn, Julien Fabre, Damian Gryski, Evan Phoenix, and Achille Roussel<br></span>
</p>
<p class="blogsummary">
  Go 1.21 adds a new port targeting the WASI preview 1 syscall API
</p>
<p class="blogtitle">
  <a href="/blog/gopls-scalability" aria-describedby="blog-description">Scaling gopls for the growing Go ecosystem</a>, <span class="date">8 September 2023</span><br>
  <span class="author">Robert Findley and Alan Donovan<br></span>
</p>
<p class="blogsummary">
  As the Go ecosystem gets bigger, gopls must get smaller
</p>
<p class="blogtitle">
  <a href="/blog/pgo" aria-describedby="blog-description">Profile-guided optimization in Go 1.21</a>, <span class="date">5 September 2023</span><br>
  <span class="author">Michael Pratt<br></span>
</p>
<p class="blogsummary">
  Introduction to profile-guided optimization, generally available in Go 1.21.
</p>
<p class="blogtitle">
  <a href="/blog/rebuild" aria-describedby="blog-description">Perfectly Reproducible, Verified Go Toolchains</a>, <span class="date">28 August 2023</span><br>
  <span class="author">Russ Cox<br></span>
</p>
<p class="blogsummary">
  Go 1.21 is the first perfectly reproducible Go toolchain.
</p>
<p class="blogtitle">
  <a href="/blog/slog" aria-describedby="blog-description">Structured Logging with slog</a>, <span class="date">22 August 2023</span><br>
  <span class="author">Jonathan Amsterdam<br></span>
</p>
<p class="blogsummary">
  The Go 1.21 standard library includes a new structured logging package, log/slog.
</p>
<p class="blogtitle">
  <a href="/blog/toolchain" aria-describedby="blog-description">Forward Compatibility and Toolchain Management in Go 1.21</a>, <span class="date">14 August 2023</span><br>
  <span class="author">Russ Cox<br></span>
</p>
<p class="blogsummary">
  Go 1.21 manages Go toolchains like any other dependency; you will never need to manually download and install a Go toolchain again.
</p>
<p class="blogtitle">
  <a href="/blog/compat" aria-describedby="blog-description">Backward Compatibility, Go 1.21, and Go 2</a>, <span class="date">14 August 2023</span><br>
  <span class="author">Russ Cox<br></span>
</p>
<p class="blogsummary">
  Go 1.21 expands Go&#39;s commitment to backward compatibility, so that every new Go toolchain is the best possible implementation of older toolchain semantics as well.
</p>
<p class="blogtitle">
<a href="/blog/all" aria-label="More articles" aria-describedby="blog-description">More articles...</a>
</p>
<div class="screen-reader-only" id="blog-description" hidden>
    Opens in new window.
</div>

    </div>

    

  </div>
</div>

<script src="/js/jquery.js"></script>
<script src="/js/playground.js"></script>
<script src="/js/play.js"></script>
<script src="/js/godocs.js"></script>

</main>
<footer class="Site-footer">
  <div class="Footer">
    <div class="Container">
      <div class="Footer-links">
          <div class="Footer-linkColumn">
            <a href="/solutions/" class="Footer-link Footer-link--primary" aria-describedby="footer-description">
              Why Go
            </a>
              <a href="/solutions/use-cases" class="Footer-link" aria-describedby="footer-description">
                Use Cases
              </a>
              <a href="/solutions/case-studies" class="Footer-link" aria-describedby="footer-description">
                Case Studies
              </a>
          </div>
          <div class="Footer-linkColumn">
            <a href="/learn/" class="Footer-link Footer-link--primary" aria-describedby="footer-description">
              Get Started
            </a>
              <a href="/play" class="Footer-link" aria-describedby="footer-description">
                Playground
              </a>
              <a href="/tour/" class="Footer-link" aria-describedby="footer-description">
                Tour
              </a>
              <a href="https://stackoverflow.com/questions/tagged/go?tab=Newest" class="Footer-link" aria-describedby="footer-description">
                Stack Overflow
              </a>
              <a href="/help/" class="Footer-link" aria-describedby="footer-description">
                Help
              </a>
          </div>
          <div class="Footer-linkColumn">
            <a href="https://pkg.go.dev" class="Footer-link Footer-link--primary" aria-describedby="footer-description">
              Packages
            </a>
              <a href="/pkg/" class="Footer-link" aria-describedby="footer-description">
                Standard Library
              </a>
              <a href="https://pkg.go.dev/about" class="Footer-link" aria-describedby="footer-description">
                About Go Packages
              </a>
          </div>
          <div class="Footer-linkColumn">
            <a href="/project" class="Footer-link Footer-link--primary" aria-describedby="footer-description">
              About
            </a>
              <a href="/dl/" class="Footer-link" aria-describedby="footer-description">
                Download
              </a>
              <a href="/blog/" class="Footer-link" aria-describedby="footer-description">
                Blog
              </a>
              <a href="https://github.com/golang/go/issues" class="Footer-link" aria-describedby="footer-description">
                Issue Tracker
              </a>
              <a href="/doc/devel/release" class="Footer-link" aria-describedby="footer-description">
                Release Notes
              </a>
              <a href="/brand" class="Footer-link" aria-describedby="footer-description">
                Brand Guidelines
              </a>
              <a href="/conduct" class="Footer-link" aria-describedby="footer-description">
                Code of Conduct
              </a>
          </div>
          <div class="Footer-linkColumn">
            <a href="https://www.twitter.com/golang" class="Footer-link Footer-link--primary" aria-describedby="footer-description">
              Connect
            </a>
              <a href="https://www.twitter.com/golang" class="Footer-link" aria-describedby="footer-description">
                Twitter
              </a>
              <a href="https://github.com/golang" class="Footer-link" aria-describedby="footer-description">
                GitHub
              </a>
              <a href="https://invite.slack.golangbridge.org/" class="Footer-link" aria-describedby="footer-description">
                Slack
              </a>
              <a href="https://reddit.com/r/golang" class="Footer-link" aria-describedby="footer-description">
                r/golang
              </a>
              <a href="https://www.meetup.com/pro/go" class="Footer-link" aria-describedby="footer-description">
                Meetup
              </a>
              <a href="https://golangweekly.com/" class="Footer-link" aria-describedby="footer-description">
                Golang Weekly
              </a>
          </div>
      </div>
    </div>
  </div>
  <div class="screen-reader-only" id="footer-description" hidden>
          Opens in new window.
  </div>
  <div class="Footer">
    <div class="Container Container--fullBleed">
      <div class="Footer-bottom">
        <img class="Footer-gopher" src="/images/gophers/pilot-bust.svg" alt="The Go Gopher">
        <ul class="Footer-listRow">
          <li class="Footer-listItem">
            <a href="/copyright" aria-describedby="footer-description">Copyright</a>
          </li>
          <li class="Footer-listItem">
            <a href="/tos" aria-describedby="footer-description">Terms of Service</a>
          </li>
          <li class="Footer-listItem">
            <a href="http://www.google.com/intl/en/policies/privacy/" aria-describedby="footer-description"
              target="_blank"
              rel="noopener">
              Privacy Policy
            </a>
            </li>
          <li class="Footer-listItem">
            <a
              href="/s/website-issue" aria-describedby="footer-description"
              target="_blank"
              rel="noopener"
              >
              Report an Issue
            </a>
          </li>
          <li class="Footer-listItem go-Footer-listItem">
            <button class="go-Button go-Button--text go-Footer-toggleTheme js-toggleTheme" aria-label="Toggle theme">
              <img
                data-value="auto"
                class="go-Icon go-Icon--inverted"
                height="24"
                width="24"
                src="/images/icons/brightness_6_gm_grey_24dp.svg"
                alt="System theme">
              <img
                data-value="dark"
                class="go-Icon go-Icon--inverted"
                height="24"
                width="24"
                src="/images/icons/brightness_2_gm_grey_24dp.svg"
                alt="Dark theme">
              <img
                data-value="light"
                class="go-Icon go-Icon--inverted"
                height="24"
                width="24"
                src="/images/icons/light_mode_gm_grey_24dp.svg"
                alt="Light theme">
            </button>
          </li>
        </ul>
        <a class="Footer-googleLogo" target="_blank" href="https://google.com" rel="noopener">
          <img class="Footer-googleLogoImg" src="/images/google-white.png" alt="Google logo">
        </a>
      </div>
    </div>
  </div>
  <script src="/js/jquery.js"></script>
  <script src="/js/carousels.js"></script>
  <script src="/js/searchBox.js"></script>
  <script src="/js/misc.js"></script>
  <script src="/js/hats.js"></script>
  <script src="/js/playground.js"></script>
  <script src="/js/godocs.js"></script>
  <script async src="/js/copypaste.js"></script>
</footer>
<section class="Cookie-notice js-cookieNotice">
  <div>go.dev uses cookies from Google to deliver and enhance the quality of its services and to 
  analyze traffic. <a target=_blank href="https://policies.google.com/technologies/cookies">Learn more.</a></div>
  <div><button class="go-Button">Okay</button></div>
</section>
</body>
</html>
`

func TestParser(t *testing.T) {
	ctx := context.Background()
	crwl := NewCrawler(nil, nil)
	data, err := crwl.Parse(ctx, bytes.NewReader([]byte(body)), nil)
	if err != nil {
		t.Fail()
		return
	}
	assert.Equal(t, 10, len(data))
	expected := []struct {
		ID   timeseriesdata.TimeSeriesDataID
		Date time.Time
	}{
		{
			ID:   timeseriesdata.TimeSeriesDataID("goblog-2023-10-09"),
			Date: time.Date(2023, time.October, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:   timeseriesdata.TimeSeriesDataID("goblog-2023-09-26"),
			Date: time.Date(2023, time.September, 26, 0, 0, 0, 0, time.UTC),
		},
	}
	for i := range expected {
		assert.Equal(t, expected[i].ID, data[i].GetID())
		assert.Equal(t, expected[i].Date, data[i].GetPublishedAt())
	}
}
