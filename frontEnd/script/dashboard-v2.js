function insertInstructor() {
  window.location.replace("createInstructor.html");
}

function setdashboard() {
  window.location.replace("dashboard-v2.html");
}
async function logout() {
  let api_error;
  let getDetails = await fetch(
    `https://dolphin-app-2zya2.ondigitalocean.app/logout?token=${getCookie(
      "token"
    )}`,
    {
      method: "GET",
    }
  ).catch((err) => {
    api_error = err;
  });
  if (api_error == "TypeError: Failed to fetch") {
    alert("Internal Server Error Please Login Again");
    window.location.replace("index.html");
    return "";
  }
  console.log(getDetails);
  window.location.replace("index.html");
}
function setViewProfile() {
  window.location.replace("viewProfile.html");
}

function insertCourse() {
  window.location.replace("createCourse.html");
}

function showCourse() {
  window.location.replace("showCourse.html");
}

function insertStudent() {
  window.location.replace("addStudent.html");
}

function showCourse() {
  window.location.replace("showCourse.html");
}

function showStudent() {
  window.location.replace("showStudents.html");
}

function showInstructor() {
  window.location.replace("instructorDetails.html");
}

let i_name = document.querySelector(".logged-user");
i_name.innerHTML = localStorage.getItem("username");

function setViewStudentProfile() {
  window.location.replace("viewStudents.html");
}
