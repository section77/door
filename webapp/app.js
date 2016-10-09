var door = angular.module('door', []);

door.controller('MainCtrl', ['$http', '$interval', function($http, $interval){
  var self = this;

  function updateState(){
    $http.get("state").then(function(res){
      self.state = res.data;
    });
    $http.get("history").then(function(res){
      self.history = res.data.slice().reverse();
    });
  };

  function callEndpointAndUpdateState(endpoint){
    $http.get(endpoint).then(function(){
      self.updateState();
    });
  }

  self.toggle = function(){
    callEndpointAndUpdateState("toggle")
  }
                             
  self.open = function(){
    callEndpointAndUpdateState("open")
  }

  self.close = function(){
    callEndpointAndUpdateState("close")
  }

  $interval(updateState, 1000);
  updateState();
}]);
