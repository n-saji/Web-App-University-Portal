function setdashboard() {
  window.location.replace("dashboard-v2.html");
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
  for (let i = 0; i < all_course_response.length; i++) {
    let each_value = all_course_response[i];
    let table1 = document.getElementById("course_table");
    let tr = document.createElement("tr");
    tr.innerHTML = `<td>${i + 1}</td>
         <td id=${i}>${each_value.course_name}</td>
         <td><button onclick=openForm(${i}) class="update_button">U</button></td>
         <td><button onclick=deletePopupTrue(${i}) class="delete_button">X</button></td>`;
    table1.appendChild(tr);
  }
}
populateCourse();

async function deleteCourse(index) {
  let index_name = document.getElementById(String(index));
  let cookie_token = getCookie("token");
  let deleteCourse = await fetch(
    `http://localhost:5050/delete-course/${index_name.innerHTML}`,
    {
      method: "DELETE",
      headers: { Token: cookie_token },
    }
  );
  let response = await deleteCourse.json();
  if (!deleteCourse.ok) {
    console.log("failed", response.Message);
    let msg = document.querySelector(".error-message");
    msg.innerHTML = "Cannot delete course! <br> Error:" + response.Message;
  } else {
    console.log("success", response);
    window.location.reload();
  }
}

async function updateCourse(course) {
  //let index_name = document.getElementById(String(index));
  let new_course_value = document.getElementById("course_name").value;
  let cookie_token = getCookie("token");
  let updateCourse = await fetch(
    `http://localhost:5050/update-course/${course}`,
    {
      method: "PATCH",
      headers: { Token: cookie_token },
      body: JSON.stringify({
        course_name: new_course_value,
      }),
    }
  );
  let response = await updateCourse.json();
  if (!updateCourse.ok) {
    console.log("failed", response);
  } else {
    console.log("success", response);
    window.location.reload();
  }
  let popup = document.getElementById("popup");
  popup.classList.remove("open-popup");
}
function openForm(index_value) {
  let popup = document.getElementById("popup");
  popup.classList.add("open-popup");
  let old_course = document.getElementById("old_course");
  let index_name = document.getElementById(String(index_value));
  old_course.innerHTML = index_name.innerHTML;
}
window.onkeydown = function (event) {
  if (event.keyCode == 27) {
    closeForm();
  }
};
function closeForm() {
  let popup = document.getElementById("popup");
  popup.classList.remove("open-popup");
}
function callUpdateFunction() {
  let old_course = document.getElementById("old_course").innerHTML;
  updateCourse(old_course);
}

function deleteTrue() {
  deleteCourse(indexForDeletePopUp1);
}