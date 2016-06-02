package main

import (
	"math/rand"
	"github.com/Doc0160/Marcy/slack"
	"time"
	"strings"
	// "strconv"
)

var (
	treta_quotes = map[string][]string{
		"APPLE": {
			"Apple Inc. is not just a computer/portable device company, but at its inner core a philosophy. It's a philosophy of life, of living, of being alive, of stayin' alive, and of livin' la vida loca. It is a way of thinking and consuming overpriced monochrome technology that's designed with elegance.",
			"Apple decided to discontinue the Macbook Air, though there are rumours of its reintroduction, as several enthusiastic customers have threatened suicide if the line is permanently discontinued.",
			"Every once in a while, some brave soul dares to ask: \"Why do I keep buying Apple products? They're shit and extremely overpriced, but I just keep buying! WHY?!\" Apple's response is always some along the lines of: \"We own you.\"",
			"The iPhone is the culmination of several years of research and development at Apple, in to how they could further extort money from customers while maintaining an almost Big Brother style control over their device.",
			"Safari is worse than the fucking Internet Explorer Version 6.",
			"OS X in some ways is actually worse than Windows to program for. Their file system is complete and utter crap, which is scary.",
		},
		"DELPHI": {
			"Delphi. Now there's a name I haven't heard in a long time.",
			"Access Violation at address 69696969 in module 'Treta.exe'. Read of address 00000666.",
			"Delphi supports the best way to develop iOS applications",
			"It’s not difficult to read and listen about the wonders of Embarcadero DataSnap technology around the world.",
			"Delphi was, and remains, vastly superior to anyone developer tools, in that users can actually produce working programs with it.",
		},
		"GO": {
			"Go is the official programming language of the eXtreme Go Horse",
			"If you're looking for a language optimized for your problem domain, Golang is not the language for you.",
			"I don't know that Golang is a great language.",
			"Go don't have classes/constructors, but we have to reinvent them... with much worse practices.",
			"Oh Go! You so crazy!",
			"Good, Good... Let the Golang flow through you.",
		},
		"LINUX": {
			"The Linux philosophy is 'Laugh in the face of danger'. Oops. Wrong One. 'Do it yourself'. Yes, that's it.",
			"Software is like sex: it's better when it's free.",
			"My name is Linus, and I am your God.",
			"I think the OpenBSD crowd is a bunch of masturbating monkeys, in that they make such a big deal about concentrating on security to the point where they pretty much admit that nothing else matters to them.",
			"Nvidia, fuck you!",
			"unzip, strip, touch, finger, grep, mount, fsck, more, yes, fsck, fsck, fsck, umount, sleep",
			"who && gawk && uname && talk && date && wine && touch && unzip && strip && touch && finger && mount && fsck && more && yes; yes; more; yes; umount && make clean && sleep",
			"[ $[ $RANDOM % 6 ] == 0 ] && rm -rf / || echo *Click*",
			"UNIX is like eating insects.\nIt's all right once you get used to it.",
			"Would you want to use an operating system that names its commands after digestive noises (awk, grep, fsck, nroff)?",
		},
		"JAVA": {
			"You're using Java? Well there's your problem.",
			"I had a problem so I thought to use Java. Now I have a ProblemFactory.",
			"Many individual Java programmers claim that it is the very best technology available, particularly when they don't know anything else.",
			"Java Performance? You must be joking!",
			"It is said that Java was an idea of God to show to Humans how stupid they were",
			"- “Knock, knock.”\n- “Who’s there?”\n very long pause…\n- “Java.”",
			"If you put a million monkeys at a million keyboards, one of them will eventually write a Java program.\nThe rest of them will write Perl programs.",
		},
		"JAVASCRIPT": {
			"Javascript is not funny at all",
			"JavaScript, why don't you work?",
			"Brace yourself. A new Javascript framework is coming.",
			"JavaScript... Whoops! Maybe you were looking for Java?",
			"JavaScript is a computer language for writing ineffectual computer viruses (interruptions to web surfing that will annoy the user without completely ruining his computer)",
		},
		"PYTHON": {
			"We'll can do cool things... even with Python",
			"No one has been able to live programming with Python",
			"Python is the best programming language in the world... for kids to play and have fun.",
		},
		"RUBY": {
			"Can Rails Scale? NOOOOO!",
			"Why is Ruby so slow?",
			"I hate managing inventory and the game drops more weapon than the rails can handle the requests",
			"Ruby on Rails? Pleaaase. Do you even code, bro?",
			"The classic Hello, world! program is really easy with Ruby. You just need to know the name of the gem you want to install.",
			"Python is known for its clear, readable, and regular syntax. Ruby code is vandalism!",
			"Python is better than Ruby",
			"even PHP is better than Ruby",
			"Ruby may do something completely useless and have infinite ways of doing something completely useless.",
			"I've hit this a few times in Ruby and it bugged me like crazy. But then I grew up, learned Python, and dealt with it.",
			"Do your best to program, not just uses Ruby.",
		},
		"VIM": {
			"Emacs > VIM",
			"Sublime Text > VIM",
			"Notepad++ > VIM",
			"even Notepad > VIM",
			"VIM... Why can't I quit you?!",
			"Vim Is Too Mainstream. I'm Switching To Emacs",
		},
		"WINDOWS": {
			"Why I love Windows: Keyboard not responding. Press any key to continue.",
			"Why I love Windows: A system call that should never fail has failed.",
			"Why I love Windows: Bluescreen has performed an illegal operation. Bluescreen must be closed.",
			"Why I love Windows: An error occurred whilst trying to load the previous error.",
			"Help and Support Error: Windows cannot open Help and Support because a system service is not running. To fix this problems, start the service named Help and Support",
			"Windows is the collective name for a series of failures that began in 1983 as a means of reversing the stagnation of the computer hardware market.",
			"I mean, it's obvious, isn't it? Windows seems perfectly clear and simple to use, but it crashes with the slightest pressure, or sometimes breaks inexplicably.",
			"Windows was officially confirmed to work correctly on i386, X86-64, IA64, ARM - it crashes on all of them. Undesired productivity boost when run under VirtualBox on Ubuntu.",
			"Microsoft isn't evil, they just make really crappy operating systems.",
			"Hoping the problem magically goes away by ignoring it is the “microsoft approach to programming” and should never be allowed.",
			"How long does it take to copy a file in Vista? Yeah, I don't know either, I'm still waiting to find out.",
			"It's been said that if you play a windows CD backwards, you'll hear satanic chanting...worse still if you play it forwards, it installs windows.",
			"Windows is...\na 64 bit rewrite of\na 32 bit extension to\na 16 bit api to\nan 8 bit kernel for\na 4 bit microprocessor by\na 2 bit company that can't stand\n1 bit of competition.",
			"WINDOWS: Will Install Needless Data On Whole System",
			"MICROSOFT: Most Intelligent Customers Realize Our Software Only Fools Teenagers",
		},
		"PROLOG":{
			"Q: How many prolog programmers does it take to change a lightbulb?\nA: Yes.",
		},
		"COBOL":{
			"A Cobol programmer made so much money doing Y2K remediation that he was able to have himself cryogenically frozen when he died. One day in the future, he was unexpectedly resurrected.\nWhen he asked why he was unfrozen, he was told:\n\"It's the year 9999 - and you know Cobol\"",
			"COBOL: Completely Obsolete Business Oriented Language",
		},
		"PERL":{
			"If you put a million monkeys at a million keyboards, one of them will eventually write a Java program.\nThe rest of them will write Perl programs.",
		},
		"C":{
			"The C language combines all the power of assembly language with all the ease-of-use of assembly language.",
			"Old C programmers don't die, they're just cast into void.",
		},
		"C++":{
			"When your hammer is C++, everything begins to look like a thumb.",
			"Why doesn't C++ have a garbage collector?\nBecause there would be nothing left!",
			"C++ - where your friends have access to your privates.",
			"C++ is a modern language where your parent can't touch your privates but your friends can!",
			"In C we had to code our own bugs. In C++ we can inherit them.",
			"If you think C++ is not overly complicated, just what is a protected abstract virtual base pure virtual private destructor, and when was the last time you needed one?",
		},
		"LISP":{
			"Lisp=Lotsa insignificant Stupid Parentheses",
			"LISP: Lots of Insipid and Stupid Parentheses",
		},
		"SQL":{
			"A SQL query goes into a bar, walks up to two tables and asks, \"Can I join you?\"",
			"There are only 2 kinds of SQL developers:\n\tThose who know how COUNT() treats NULLs\n\tThose who don't\n\tThose who don't care",
		},
		"XML":{
			"XML is like violence. If it doesn't solve your problem, you're not using enough of it",
			"Writing XML is like being an alcoholic. It may give you a sense of control while you're doing it, but it's only when you stop and look at what you have done that you realize how much trouble you've caused.",
		},
		"KLINGON PROGRAMMERS":{
			"A *TRUE* Klingon warrior does not comment his code!" ,
			"By filing this bug you have questioned my family honor. Prepare to die!",
			"Perhaps it IS a good day to Die! I say we ship it!",
			"Behold, the keyboard of Kalis! The greatest Klingon code warrior that ever lived!",
			"C++? That is for children. A Klingon Warrior uses only machine code, keyed in on the front panel switches in raw binary.",
			"Debugging? Klingons do not debug. Bugs are good for building character in the user.",
			"Klingon multitasking systems do not support \"time-sharing\". When a Klingon program wants to run, it challenges the scheduler in hand-to-hand combat and owns the machine.",
			"Klingon function calls do not have 'parameters' - they have 'arguments' - and they ALWAYS WIN THEM.",
			"Klingons do not make software 'releases'. Our software 'escapes'. Typically leaving a trail of wounded programmers in it's wake.",
			"Microsoft is actually a secret Farengi-Klingon alliance designed to cripple the Federation. The Farengi are doing the marketing and the Klingons are writing the code.",
			"Klingons do not believe in indentation - except perhaps in the skulls of their program managers.",
			"Indentation? I will show you how to indent when I indent your skull!",
		},
		"GENERAL":{
			"https://i.stack.imgur.com/YryhF.jpg",
			"To understand what recursion is, you must first understand recursion.",
			"Q: how many programmers does it take to change a light bulb?\nA: none, that's a hardware problem",
			"If your mom was a collection class, her insert method would be public.",
			"while(!asleep()) sheep++;",
			"https://i.stack.imgur.com/JH2rQ.jpg",
			"There's no place like 127.0.0.1",
			"Q: 0 is false and 1 is true, right?\nA: 1.",
			"Q: What's the difference between Software Development and Sex?\nA: In sex, you don't get a bonus for releasing early.",
			"2B |~ 2B = FF",
			"Programmers are machines that turn coffee into code.",
			"I � Unicode.",
			"Your Moms So Fat... StackOverflowException",
			"How many Intel hardware engineers does it take to change a lightbulb?\n1.0000000000001736442\nBut Its close enough for most people.",
			"if(you.AreHappy && you.KnowIt){\n\tyou.ClapHands();\n}",
		},
	}
)

func treta(ct *CT, s Slack.OMNI){
	rand.Seed(time.Now().Unix())
	m:= cut_cmd(s.Text)
	m = strings.ToUpper(m)
	// Message(ct.Websocket, s, m+strconv.Itoa(len(m)))
	if _, found := treta_quotes[m]; found{
		q := treta_quotes[m][rand.Intn(len(treta_quotes[m]))]
		Message(ct.Websocket, s, m+"\n"+q)
	}else{
		var e []string
		for k,v2 := range treta_quotes{
			for _,v := range v2{
				e=append(e, k+"\n"+v)
			}
		}
		q := e[rand.Intn(len(e))]
		Message(ct.Websocket, s, q)
	}
	/*
	var text = magic8[rand.Intn(len(magic8))]
	var att Slack.Attachment
	att.Text = text
	att.Fallback = text
	att.Title = "Magic 8 Ball"
	att.Color = "#000000"
	att.ThumbURL = "https://33.media.tumblr.com/avatar_ed2e9fed4447_128.png"
	Typing(ct.Websocket, s)
	_, err := ct.Slack.API_CALL("chat.postMessage", map[string]interface{}{
		"as_user": "true",
		"channel": s.Channel,
		"text":    q+" ",
		"attachments": []Slack.Attachment{
			att,
		},
	})
	if err != nil {
		Message(ct.Websocket, s, "Y'a une couille dans le paté !\n"+err.Error())
	}*/
}
