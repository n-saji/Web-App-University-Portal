function removeError() {
  let instructorcode = document.getElementById("ic");
  let instructorname = document.getElementById("in");
  let department = document.getElementById("dp");
  let coursename = document.getElementById("cn_drop_down");
  instructorcode.classList.remove("error");
  instructorname.classList.remove("error");
  department.classList.remove("error");
  coursename.classList.remove("error");
}
async function InsertInstructorValues() {
  let instructorcode = document.getElementById("ic");
  let instructorname = document.getElementById("in");
  let department = document.getElementById("dp");
  let coursename = document.getElementById("cn_drop_down");
  let error_log = document.getElementById("error_log");
  if (instructorcode.value === "") {
    instructorcode.classList.add("error");
  }

  if (instructorname.value === "") {
    instructorname.classList.add("error");
  }

  if (department.value === "") {
    department.classList.add("error");
  }

  if (coursename.value == "Choose Course") {
    coursename.classList.add("error");
  }
  let createInstructor = await fetch(
    `http://localhost:5050/insert-instructor-details`,
    {
      method: "POST",

      body: JSON.stringify({
        InstructorCode: instructorcode.value,
        InstructorName: instructorname.value,
        Department: department.value,
        CourseName: coursename.value,
      }),
    }
  );
  let response = await createInstructor.json();
  console.log(response);
  if (createInstructor.ok != true) {
    error_log.innerHTML = "Error submitting!";

  } else if (createInstructor.ok == true) {
    document.getElementById("responseBody").innerHTML =
      "Added<br> Please Create EmailId and Password";

    let rtl = document.getElementById("redirect_to_login");
    rtl.innerHTML = "Create Account";
    let URL = `http://localhost:5050` + response.URl;
    document.cookie = `url=${URL}`;
    localStorage.setItem("URL_Create_Login", URL);
  }
}

async function create_account_new() {
  let instructorcode = document.getElementById("ic").value;
  let instructorname = document.getElementById("in").value;
  let department = document.getElementById("dp").value;
  let coursename = document.getElementById("cn_drop_down").value;
  if (
    !instructorcode ||
    !instructorname ||
    !department ||
    coursename == "Choose Course"
  ) {
    alert("Please fill details");
    return;
  }
  window.location.replace("createInstructorAccount.html");
}
function setbackpage() {
  window.location.replace("instructorDetails.html");
}

function setdashboard() {
  window.location.replace("allinstructor.html");
}
