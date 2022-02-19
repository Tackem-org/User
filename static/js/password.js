(function () {
    $(() => {
        $('form').submit(function () {
            let $op = $('[name="op"]')
            if ($op.val() == '') {
                alert("Original Password Field Blank")
                return false;
            }
            let $np1 = $('[name="np1"]')
            if ($np1.val() == '') {
                alert("New Password Field Blank")
                return false;
            }
            let $np2 = $('[name="np2"]')
            if ($np2.val() == '') {
                alert("New Password Field Blank")
                return false;
            }
            if ($np1.val() != $np2.val()) {
                alert("Passwords Don't Match")
                return false;
            }
            return true;
        });
    });

})();
