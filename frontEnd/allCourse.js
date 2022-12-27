let dropdown = document.getElementById("cn_drop_down");
dropdown.length = 0;

let defaultOption = document.createElement("option");
defaultOption.text = "Choose Course";

dropdown.add(defaultOption);
dropdown.selectedIndex = 0;

const url = 'http://localhost:5050/retrieve-all-courses';

fetch(url)
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
        option.text = data[i].CourseName;
        dropdown.add(option);
      }
    });
  })
  .catch(function (err) {
    console.error("Fetch Error -", err);
  });
