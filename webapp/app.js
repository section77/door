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

  function callEndpointAndUpdateState(endpoint){
    $http.get(endpoint).then(function(){
      self.updateState();
    });
  }

  self.toggle = function(){
    callEndpointAndUpdateState("/toggle")
  }
                             
  self.open = function(){
    callEndpointAndUpdateState("/open")
  }

  self.close = function(){
    callEndpointAndUpdateState("/close")
  }

  $interval(update, 1000);
  update();
}]);
