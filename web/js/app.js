var Router = Backbone.Router.extend({
  routes: {
    '/': 'home'
  },

  initialize: function() {
    Backbone.history.start();
  },

  home: function(){
//    var view = new ProjectListView();
//    view.render();
  }
});

//HI