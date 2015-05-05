$(document).ready(function(){
  $(window).scroll(function(){
    var wScroll = $(this).scrollTop();
console.log(wScroll, wScroll * 3, wScroll * 1.5)
    $('.rain').css({
      'transform' : 'translate(0px, '+ wScroll /8 +'%)'
    });

    $('.landing-footer').css({
      'transform' : 'translate(0px, '+ wScroll *1.25 +'%)'
    });

  });

});
