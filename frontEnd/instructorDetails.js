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
    tr.innerHTML = `<td>${each_value.instructor_code}</td>
       <td id=${i}>${each_value.instructor_name}</td>
       <td>${each_value.department}</td>

       <td id=${each_value.course_name + i}>${each_value.course_name}</td>
       <td><button class="update_button">U</button></td>
       <td><button onclick=deleteInstructor(${i},${
      each_value.course_name + i
    }) class="delete_button">X</button></td>`;

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

async function deleteInstructor(index, course_name_index) {
  console.log(index,course_name_index)
  let index_name = document.getElementById(index);
  console.log(index_name, course_name_index);
  let cookie_token = getCookie("token");
  let deleteCourse = await fetch(`http://localhost:5050/delete-instructor`, {
    method: "DELETE",
    headers: { Token: cookie_token },
    body: JSON.stringify({
      instructor_name: index_name.innerText,
      course_name: course_name_index.innerText,
    }),
  });

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
