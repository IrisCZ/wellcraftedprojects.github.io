var opened = false;
var newOpened = false;
var loginOpened = false;
$(document).ready(function(){
  $(".nav-icon").on('click', function() {
          if (!opened){
          	if (!newOpened && !loginOpened){
          		$(this).toggleClass("nav-icon-close");
          	}
          	if(newOpened){
          		document.getElementById("new-project").close();
          		newOpened = false;
          	}
          	if(loginOpened){
          		document.getElementById("login-window").close();
          		loginOpened = false;
          	}
          	
            document.getElementById("summary").show();        		
          } else {
          	$(this).toggleClass("nav-icon-close");
            document.getElementById("summary").close();
          }
          opened = !opened;
        });
	    
        $('#add-project').click(function(){
          if (!newOpened){
            document.getElementById("summary").close();
            document.getElementById("new-project").show();
            newOpened = true;
            opened = false;

          }
        });
        
        $('#login').click(function(){
          if (!loginOpened){
            document.getElementById("summary").close();
            document.getElementById("login-window").show();
            loginOpened = true;
            opened = false;

          }
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