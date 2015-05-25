var opened = false;
var newOpened = false;
var loginOpened = false;
$(document).ready(function(){
  $(".nav-icon").on('click', function() {
  	if (opened){
  		if (newOpened){
    	  document.getElementById("new-project").close();
    	  newOpened = false;
    	  $(".nav-icon-arrow").hide();
  		} else if (loginOpened){
    	  document.getElementById("login-window").close();
    	  loginOpened = true;
    	  $(".nav-icon-arrow").hide();
  		} else {
		  document.getElementById("summary").close();
  		}
  	} else {
  		document.getElementById("summary").show(); 		  		
  	}
  	$(this).toggleClass("nav-icon-close");
  	opened = !opened;
  });
	    
  $('#add-project').click(function(){
      document.getElementById("summary").close();
      document.getElementById("new-project").show();
      newOpened = true;
      $(".nav-icon-arrow").show();
  });

  $('.nav-icon-arrow').click(function () {
    $(".nav-icon-arrow").hide();
    if (newOpened){
      document.getElementById("new-project").close();
      newOpened = false;
  	} else if (loginOpened){
      document.getElementById("login-window").close();
      loginOpened = false;
	}
	document.getElementById("summary").show();
	opened = true;   	
  });
        
  $('#login').click(function(){
    document.getElementById("summary").close();
    document.getElementById("login-window").show();
    loginOpened = true;
    $(".nav-icon-arrow").show();
  });

  $('.summary-list a:not(#add-project, #login)').click(function(){
   	$('#toggle-modal').click();
  });
  $('.project').each(function(index){
   	var images = ['logoRunU-bueno','spring','go','grails'];
   	var image = images[index] || 'grails';
  	$(this).find('.project-img').css({
   		'background-image':'url(images/projects/'+image+'.png)'
   	})
  });
});