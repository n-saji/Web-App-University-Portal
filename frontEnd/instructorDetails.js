async function populateInstructors() {
  let cookie_token = getCookie("token");
  let all_instructors = await fetch(
    `http://localhost:5050/retrieve-instructors`,
    {
      credentials: "same-origin",
      headers: { token: cookie_token },
    }
  );
  let all_instructors_response = await all_instructors.json();
  if (all_instructors_response == "authentication time-out") {
    alert("Timed-out re login");
    setTimeout(window.location.replace("index.html"), 2000);
  }
  for (let i = 0; i < all_instructors_response.length; i++) {
    let each_value = all_instructors_response[i];
    let table1 = document.getElementById("instructor_table");
    let tr = document.createElement("tr");
    tr.innerHTML = `<td>${each_value.InstructorCode}</td>
       <td id=${i}>${each_value.InstructorName}</td>
       <td>${each_value.Department}</td>
       <td>${each_value.CourseName}</td>
       <td><button>U</button></td>
       <td><button onclick=deleteInstructor(${i}) class="delete_button">X</button></td>`;
    table1.appendChild(tr);
  }
}
populateInstructors();
function setdashboard() {
  window.location.replace("allinstructor.html");
}
function setbackpage() {
  window.location.replace("createInstructor.html");
}
async function deleteInstructor(index) {
  let index_name = document.getElementById(String(index));

  //let response_for_deleteion = document.getElementById("response_for_deleteion");

  let deleteCourse = await fetch(
    `http://localhost:5050/delete-instructor/${index_name.innerHTML}`,
    {
      method: "DELETE",
      credentials: "same-origin",
    }
  );
  let response = await deleteCourse.json();
  if (!deleteCourse.ok) {
    console.log("failed", response);
  } else {
    console.log("success", response);
    window.location.reload();
  }
}
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
