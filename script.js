const colorItems = document.querySelectorAll('.color-item');
const square = document.getElementById('square');

colorItems.forEach(item => {
    item.addEventListener('click', () => {
        square.style.backgroundColor = item.classList[1];
    });
});
