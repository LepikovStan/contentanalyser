!function(e) {
    let data = {
        referrer: "",
        title: "",
        url: "",
        author: "",
        id: "",
        event: "",
        language: "",
        time_on_page: 0,
        scroll_to_middle: false,
        scroll_to_end: false,
        device: "",
        user_agent: navigator.userAgent,
        tags: "",
        keywords: "",
    };
   data.href = window.location.href;
   data.referrer = document.referrer;
   data.title =  document.querySelector("meta[property='og:title']")?.getAttribute("content");
   data.url =  document.querySelector("meta[property='og:url']")?.getAttribute("content");
   data.author = document.querySelector("meta[property='og:author']")?.getAttribute("content");
    let authorTag = document.querySelector("div[itemprop='author']");
    if ((authorTag !== undefined && authorTag !== null) && data.author === undefined) {
        data.author = authorTag.querySelector("meta[itemprop='name']")?.getAttribute("content");
    }

    function send(d) {
        let xhr = new XMLHttpRequest();
        xhr.open("POST", "https://d5d89238ickq1pv2dk0f.apigw.yandexcloud.net//upload");

        xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            console.log(xhr.status);
            console.log(xhr.responseText);
        }};


        xhr.send(JSON.stringify(d));
    }


   const fpPromise = import('https://openfpcdn.io/fingerprintjs/v3')
   .then(FingerprintJS => FingerprintJS.load())

 // Get the visitor identifier when you need it.
 fpPromise
   .then(fp => fp.get())
   .then(result => {
     // This is the visitor identifier:
        data.id = result.visitorId
        data.event = "load"
        console.log(data);

       send(data)
   })



  var body = document.body,
  html = document.documentElement;

var height = Math.max( body.scrollHeight, body.offsetHeight, 
                     html.clientHeight, html.scrollHeight, html.offsetHeight );

    var sendMiddlePage = (function(){
        var done = false;
        return function() {
            if (!done) {
                done = true;

                data.event = "scroll_middle"
                data.scroll_to_middle = true
                send(data)
            }
        }
    })()

    var sendBottomPage = (function(){
        var done = false;
        return function() {
            if (!done) {
                done = true;

                data.event = "scroll_bottom"
                data.scroll_to_end = true
                send(data)
            }
        }
    })()

   window.addEventListener('scroll', function() {
    if (window.scrollY > height/2) {
        sendMiddlePage()
    }
    if (window.scrollY +this.window.innerHeight > height-100) {
        sendBottomPage()
    }
  });

}();