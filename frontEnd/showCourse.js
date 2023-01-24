function setdashboard() {
  window.location.replace("allInstructor.html");
}
function setaddcourse() {
  window.location.replace("createCourse.html");
}
async function populateCourse() {
  let cookie_token = getCookie("token");
  let all_course = await fetch(`http://localhost:5050/retrieve-all-courses`, {
    credentials: "same-origin",
    headers: { token: cookie_token },
  });
  let all_course_response = await all_course.json();
  if (all_course_response == "authentication time-out") {
    alert("Timed-out re login");
    setTimeout(window.location.replace("index.html"), 2000);
  }
  for (let i = 0; i < all_course_response.length; i++) {
    let each_value = all_course_response[i];
    let table1 = document.getElementById("course_table");
    let tr = document.createElement("tr");
    tr.innerHTML = `<td>${i + 1}</td>
         <td id=${i}>${each_value.course_name}</td>
         <td><button class="update_button">U</button></td>
         <td><button onclick=deleteCourse(${i}) class="delete_button">X</button></td>`;
    table1.appendChild(tr);
  }
}
populateCourse();
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
async function deleteCourse(index) {
  let index_name = document.getElementById(String(index));
  console.log(index_name.innerHTML);
  let deleteCourse = await fetch(
    `http://localhost:5050/delete-course/${index_name.innerHTML}`,
    {
      method: "DELETE",
      headers: { Token: cookie_token },
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
