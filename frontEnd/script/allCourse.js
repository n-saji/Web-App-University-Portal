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
