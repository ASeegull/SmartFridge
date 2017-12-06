(function() {
    //Creates constants for cashed elements
    const GetContentBtn = document.getElementById('getContents');
    const GetRecipesBtn = document.getElementById('getRecipes');

    //Creates host constant
    const host = 'http://localhost:8080';

    //Adds eventlisteners to buttons
    GetContentBtn.addEventListener('click', getContent);
    GetRecipesBtn.addEventListener('click', getRecipes);

    //Abstract GET request
    function fetchData(query) {
        return fetch(query, {
            method: "GET",
            headers: {
                'Accept': 'application/json, text/plain, */*',
                "Content-Type": "application/json, text/plain"
            }
        }).then(res => {
            let contentType = res.headers.get("content-type");
            console.log(contentType);
            if(contentType && contentType.indexOf("application/json") !== -1) {
              return res.json().then(data => {
                console.log(data);
              });
            } else {
              return res.text().then(text => {
                console.log(text);
                document.write(text);
              });
            }})
          .catch(e => console.log(e));
        }

    //Fetches content of user's fridge from server
    function getContent() {
        let jsn = fetchData('/content');
    }

    //Get's recipes based on content of fridge
    function getRecipes() {
        let jsn = fetchData('/recipes');
    }

    const modal = {
        "signup": document.getElementById('signup'),
        "login": document.getElementById('login')
    }

    let closeBtn = document.querySelectorAll('.closeBtn');
    closeBtn.forEach(btn => btn.addEventListener('click', closeModal));

    let modalBtn = document.querySelectorAll('.btn-nav');
    modalBtn.forEach(btn => btn.addEventListener('click', openModal));

    document.addEventListener('click', clickOutside);

    function openModal(e) {
        let query = e.currentTarget.dataset.modal;
        let show = modal[query];
        show.classList.remove('modal-hidden');
        show.classList.add('modal-show');
    }

    function closeModal(e) {
        let query = e.currentTarget.dataset.modal;
        let hide = modal[query];
        hide.classList.remove('modal-show');
        hide.classList.add('modal-hidden');
    }

    function clickOutside(e) {
        if (e.target == modal) {
            closeModal();
        }  
    }


    const signUpForm = document.getElementById('signup');
    
    signUpForm.addEventListener('submit', submit)
    
    function submit(e) {
        e.preventDefault();

        let userdata = [];
        let fields = signUpForm.querySelectorAll('input');
        console.log(fields);
        fields.forEach(field => userdata.push(field.value))
        let data = {
            name: userdata[0], 
            password: userdata[1],
            email: userdata[3]
        }

        console.log(data);
        fetch("/signup", {
            method: 'post',
            body: JSON.stringify(data)
          }).then(res => window.location.href = res.url)
            .catch(err => console.log(err))
    }
})();

