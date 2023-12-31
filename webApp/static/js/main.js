$(document).ready(function() {
    $('#catSelect').on('mouseenter', 'option', function() {
        $(this).css('background-color', '#D7AF70');
    });
    $('#catSelect').on('mouseleave', 'option', function() {
        $(this).css('background-color', '#383b37');
    });
});