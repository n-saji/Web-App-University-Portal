async function InsertInstructorValues() {
  let instructorcode = document.getElementById("ic").value;
  if (!instructorcode) {
    alert("Please enter instructorcode.");
    return;
  }
  let instructorname = document.getElementById("in").value;
  if (!instructorname) {
    alert("Please enter instructorname.");
    return;
  }
  let department = document.getElementById("dp").value;
  if (!department) {
    alert("Please enter department.");
    return;
  }
  let coursename = document.getElementById("cn_drop_down").value;
  console.log(coursename);
  if (coursename == "Choose Course") {
    alert("Please choose a course.");
    return;
  }
  let createInstructor = await fetch(
    `http://localhost:5050/insert-instructor-details`,
    {
      method: "POST",

      body: JSON.stringify({
        InstructorCode: instructorcode,
        InstructorName: instructorname,
        Department: department,
        CourseName: coursename,
      }),
    }
  );
  let response = await createInstructor.json();
  console.log(response);
  if (createInstructor.ok != true) {
    alert(response.Err);
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
  window.location.replace("index.html");
}

function setdashboard() {
  window.location.replace("allinstructor.html");
}
