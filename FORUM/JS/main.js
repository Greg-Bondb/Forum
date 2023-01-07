const burger = document.querySelector('.burger')
const body = document.querySelector('body');
burger.addEventListener('click', ()=>{
    burger.classList.toggle('active')
    body.classList.toggle('open')
})

var searchInput = document.getElementById('searchInput');
var searchInputResp = document.getElementById('searchInputResp');

searchInput.addEventListener('click', function() {
    if (window.location.pathname === "/accueil") {
        window.location.replace("/subject_affichage");
    }
})

searchInput.addEventListener('keyup', function() {
    var input = searchInput.value;

    var divs = document.getElementsByClassName('range');

    for (i = 0; i < divs.length; i++) {
        if(divs[i].getAttribute('data-title').toLocaleLowerCase().includes(input.toLocaleLowerCase())) {
            divs[i].style.display = 'block';
        }
        else {
            divs[i].style.display = 'none';
        }
    }
})

searchInputResp.addEventListener('click', function() {
    if (window.location.pathname === "/accueil") {
        window.location.replace("/subject_affichage");
    }
})

searchInputResp.addEventListener('keyup', function() {
    var input = searchInputResp.value;

    var divs = document.getElementsByClassName('range');

    for (i = 0; i < divs.length; i++) {
        if(divs[i].getAttribute('data-title').toLocaleLowerCase().includes(input.toLocaleLowerCase())) {
            divs[i].style.display = 'block';
        }
        else {
            divs[i].style.display = 'none';
        }
    }
})
