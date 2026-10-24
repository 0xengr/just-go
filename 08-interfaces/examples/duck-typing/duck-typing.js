// NOTE: Dynamic languages like Ruby, Python, or JavaScript have no interface,
// so they use duck typing. It simply means: if it behaves like a duck and
// quacks like a duck, then it is a duck.

// NOTE: Here is the same idea in JavaScript.

function notify(notifier) {
  notifier.send("Hello");
}

const emailNotifier = {
  send(message) {
    console.log("Email: ", message);
  },
};

const smsNotifier = {
  send(message) {
    console.log("SMS: ", message);
  },
};

function main() {
  notify(emailNotifier); // Email:  Hello
  notify(smsNotifier); // SMS:  Hello
}

main();
