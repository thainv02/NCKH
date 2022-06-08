var listCoursesBlock = document.querySelector('#list-course')

var courseAPI = ''


function start(){
    getCourses(renderCourses);
}

start(); 

function getCourses(callback){
    fetch(courseAPI)
        .then(function(response){
            return response.json();
        })
        .then(callback);
}
function renderCourses(courses){
    var listCoursesBlock = document.querySelector('#list-course')
    var htmls = courses.map(function(course) {
        return
            <li>
                <h5>${course.name}</h5>
                <p>${course.description}</p>
            </li>
        ;
    });
    listCoursesBlock.innerHTML = htmls.join('');

}