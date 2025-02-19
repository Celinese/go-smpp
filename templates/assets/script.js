// Starter JavaScript for disabling form submissions if there are invalid fields
(function () {
    'use strict';
    window.addEventListener('load', function () {
        // Fetch all the forms we want to apply custom Bootstrap validation styles to
        var forms = document.getElementsByClassName('needs-validation');
        /*   var tel = document.querySelector('input[type="tel"]'); */

        // Loop over them and prevent submission
        var validation = Array.prototype.filter.call(forms, function (form) {
            form.addEventListener('submit', function (event) {
                if (form.checkValidity() === false) {
                    event.preventDefault();
                    event.stopPropagation();
                }
                /*  if (!check(tel.value)) {
                     event.preventDefault();
                     event.stopPropagation();
                     alert('error on phone')
                 } */
                form.classList.add('was-validated');
            }, false);
        });
    }, false);


})();

function check(input) {
    /* if (input.length != 10)
        return false;
    return true; */
    var tel = /^\+?([0-9]{2})\)?[-. ]?([0-9]{4})[-. ]?([0-9]{4})$/;
    if (input.match(tel)) {
        return true;
    } else {
        return false;
    }
}

function DataTB() {
    var table = $('#example').DataTable({
        lengthChange: false,
        buttons: ['copy', 'excel', 'pdf', 'colvis']
    });

    table.buttons().container()
        .appendTo('#example_wrapper .col-md-6:eq(0)');
};
