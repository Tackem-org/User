(function () {
    $(() => {
        $('form').submit(function () {
            //TODO MAKE THIS DO ALL THE POSSIBLE CHECKS FOR THE PASSWORD FIELDS
            //MAKING SURE THEY ARE NOT BLANK AND THE NEW MATCH BEFORE PASSING ONTO THE NEXT SECTION
            let $op = $('[name="op"')
            let $np1 = $('[name="np1"')
            let $np2 = $('[name="np2"')
            if ($np1.val() != $np1.val()) {
                alert("Passwords Don't Match")
                return false;
            }
            return true; // return false to cancel form action
        });
    });

})();
