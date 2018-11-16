window.onload = function() {
  var button = document.getElementById('hello');
  button.onclick = function() {
    document.getElementById('content').textContent = "Hello World!";
  };
};