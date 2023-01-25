let dropdown = document.getElementById("cn_drop_down");
let defaultOption = document.createElement("option");
let cookie_token = getCookie("token");
const url = "http://localhost:5050/retrieve-all-courses";
dropdown.length = 0;
defaultOption.text = "Choose Course";
dropdown.add(defaultOption);
dropdown.selectedIndex = 0;
fetch(url, {
  headers: { token: cookie_token },
})
  .then(function (response) {
    if (response.status !== 200) {
      console.warn(
        "Looks like there was a problem. Status Code: " + response.status
      );
      response.json().then(function (data) {
        alert(data);
        setTimeout(window.location.replace("index.html"), 2000);
      });
      return;
    }
    response.json().then(function (data) {
      let option;

      for (let i = 0; i < data.length; i++) {
        option = document.createElement("option");
        option.text = data[i].course_name;
        dropdown.add(option);
      }
    });
  })
  .catch(function (err) {
    console.error("Fetch Error -", err);
  });

function getCookie(name) {
  var nameEQ = name + "=";
  var ca = document.cookie.split(";");
  for (var i = 0; i < ca.length; i++) {
    var c = ca[i];
    while (c.charAt(0) == " ") c = c.substring(1, c.length);
    if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length, c.length);
  }
  return null;
}
