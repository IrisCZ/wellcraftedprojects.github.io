$(document).ready(function(){
  $(window).scroll(function(){
    var wScroll = $(this).scrollTop();
    $('.rain').css({
      'transform' : 'translate(0px, '+ wScroll /8 +'%)'
    });

    $('.landing-footer').css({
      'transform' : 'translate(0px, '+ wScroll *1.15 +'%)'
    });
    if (wScroll > 570) {
    	var height = wScroll - 565;
    	var top = 80 - height;
    	if (top < 0){
    		top = 0;
    	}
    	var paddingTop = height;
    	if (paddingTop > 80){
    		paddingTop = 80;
    	}
    	if (height > 110){
    		height = 110;
    	}
    	$('.projects-header').css({
    	    'height':height,
    	    'padding-top':paddingTop,
    	    'position':'fixed',
    	    'z-index':900,
    	    'width':'100%',
    	    'top':top
    	});

    	$('.logo').css({
    		'background-image':'url(images/logoN.svg)'
    	});$('.projects-container').css({
    		'margin-top': 105 - top
    	});
    	$('.nav-icon').addClass('nav-icon-black');
    } else {
    	$('.logo').css({
    		'background-image':'url(images/logo.svg)'
    	});
    	$('.projects-header').css({
    	    'height':'30px',
    	    'padding-top':0,
    	    'position':'relative',
    	    'top':0
    	});
    	$('.nav-icon').removeClass('nav-icon-black');
		$('.projects-container').css({
    		'margin-top': 0
    	});
    }
  });

});