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
  },

  template: _.template($('#project_template').html()),

  render: function(){

    console.log("PROJECTS")
    console.log(this.projects.toJSON())
    console.log(this.projects.at(1))
//    var projects = new Projects([
//          {image:"images/runu.png",author:'Author´s Name',description:'Project´s decription. Project´s decription. Project´s decription.',url:'https://github.com/runu/runu.github.io',positives:5,negatives:2},
//          {image:"images/new-project.jpg",author:'Author´s Name',description:'Project´s decription. Project´s decription. Project´s decription.',url:'https://github.com/runu/runu.github.io',positives:5,negatives:2},
//          {author:'Author´s Name',description:'Project´s decription. Project´s decription. Project´s decription.',url:'https://github.com/runu/runu.github.io',positives:5,negatives:2},
//    ]);
    this.$el.html(this.template({projects:this.projects.toJSON()}))
  }

});