//自动设置网站标题
function autSetTitle(){
  console.log('初始化-autSetTitle() => ' +document.title);
  var OriginTitile = document.title;
  var titleTime;
  document.addEventListener('visibilitychange', function() {
    if(document.hidden) {
      //$('[rel="icon"]').attr('href', "/failure.ico");
      //$('[rel="shortcut icon"]').attr('href', "/failure.ico");
      //document.title = '喔唷，崩溃啦！';
      document.title = ' 😂去哪里了！';
      clearTimeout(titleTime);
    } else {
      //$('[rel="icon"]').attr('href', "/favicon-32x32.ico");
      //$('[rel="shortcut icon"]').attr('href', "/favicon-32x32.ico");
      document.title = ' 😍欢迎回来！';
      titleTime = setTimeout(function() {
        document.title = OriginTitile;
      }, 2000);
    }
  });
}