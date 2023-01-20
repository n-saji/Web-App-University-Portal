let dropdown = document.getElementById("cn_drop_down");
dropdown.length = 0;

let defaultOption = document.createElement("option");
defaultOption.text = "Choose Course";
let cookie_token = getCookie("token");
dropdown.add(defaultOption);
dropdown.selectedIndex = 0;

const url = "http://localhost:5050/retrieve-all-courses";

fetch(url, {
  headers: { token: cookie_token },
})
  .then(function (response) {
    if (response.status !== 200) {
      console.warn(
        "Looks like there was a problem. Status Code: " + response.status
      );
      return;
    }

    // Examine the text in the response
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