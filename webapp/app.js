var door = angular.module('door', []);

door.controller('MainCtrl', ['$http', '$interval', function($http, $interval){
  var self = this;

  function update(){
    $http.get("/state").then(function(res){
      self.state = res.data;
    });
    $http.get("/history").then(function(res){
      self.history = res.data;
    });
  };


  self.open = function(){
    $http.get("/open").then(function(){
      self.updateState();
    });
  }

  self.close = function(){
    $http.get("/close").then(function(){
      self.updateState();
    });
  }

  $interval(update, 1000);
  update();
}]);
