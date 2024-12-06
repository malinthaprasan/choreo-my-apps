const express = require('express');
const app = express();
const port = 3000;

// Middleware to handle long response time
app.get('/long-response', (req, res) => {
    console.log('Request received. Starting the 100-secs wait...');
    setTimeout(() => {
        res.send('Response after 100 sec');
        console.log('Response sent after 100 sec');
    }, 100 * 1000); // 100 secs in milliseconds
});

// Start the server
app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
});
