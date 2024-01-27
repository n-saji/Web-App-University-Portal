function setbackpage() {
  window.location.replace("showCourse.html");
}

function setdashboard() {
  window.location.replace("dashboard-v2.html");
}

async function InsertCourseValues() {
  let course_name = document.getElementById("cn").value;
  let response_for_creation = document.getElementById("response_for_creation");
  if (!course_name) {
    alert("Please enter Course Name.");
    return;
  }
  let cookie_token = getCookie("token");
  let createCourse = await fetch(`http://3.111.149.112:5050/insert-course`, {
    method: "POST",
    headers: { Token: cookie_token },
    body: JSON.stringify({
      course_name: course_name,
    }),
  });
  let response = await createCourse.json();
  if (!createCourse.ok) {
    response_for_creation.innerHTML = response + "!!";
  } else {
    let response_for_creation = document.getElementById(
      "response_for_creation"
    );
    response_for_creation.innerHTML = "Successfully Created";
  }
}
