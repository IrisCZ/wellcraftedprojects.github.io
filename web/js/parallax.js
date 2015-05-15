$(document).ready(function(){
  $(window).scroll(function(){
    var wScroll = $(this).scrollTop();
    $('.rain').css({
      'transform' : 'translate(0px, '+ wScroll /8 +'%)'
    });

    $('.landing-footer').css({
      'transform' : 'translate(0px, '+ wScroll *1.15 +'%)'
    });
    if (wScroll > 560) {
    	var height = wScroll - 560;
    	var top = 60 - height;
    	if (top < 0){
    		top = 0;
    	}
    	var paddingTop = height;
    	if (paddingTop > 60){
    		paddingTop = 60;
    	}
    	if (height > 100){
    		height = 100;
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
    		'margin-top': 65 - top
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