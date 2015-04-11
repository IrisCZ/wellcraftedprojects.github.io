ProjectListView = Backbone.View.extend({

  projects: new Projects(),

  initialize: function(){
    this.$el = $('#projects_container');
    this.listenTo(this.projects,'sync',this.render, this);
    var collection = this.projects;
    this.projects.fetch({
      success: function(model, response) {
        collection.set(response.projects)
      }
    });
    this.addProjectView = new AddProjectView({el:$('#newproject_container')});
  },

  template: _.template($('#project_template').html()),

  render: function(){
    this.$el.html(this.template({projects:this.projects.toJSON()}));
    this.addProjectView.render();
  }

});