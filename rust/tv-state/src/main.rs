use rand::Rng;
use std::collections::VecDeque;
use std::thread;
use std::time::Duration;

#[derive(Debug)]
enum State {
    PowerOn,
    PowerOff,
}

#[derive(Debug)]
enum Event {
    PushedPowerButton,
    PushedVolumeUpButton,
    PushedVolumeDownButton,
}

#[derive(Debug)]
struct EventQueue(VecDeque<Event>);

impl EventQueue {
    pub fn new() -> Self {
        let d = VecDeque::new();
        EventQueue(d)
    }

    pub fn dequeue(&mut self) -> Option<Event> {
        self.0.pop_front()
    }

    pub fn enqueue(&mut self, event: Event) {
        self.0.push_back(event);
    }
}

#[derive(Debug)]
struct TV {
    state: State,
    events: EventQueue,
    volume: u8,
}

impl TV {
    pub fn new() -> Self {
        Self {
            state: State::PowerOff,
            events: EventQueue::new(),
            volume: 10,
        }
    }

    pub fn be_pushed_power_button(&mut self) {
        self.events.enqueue(Event::PushedPowerButton);
    }

    pub fn be_pushed_volume_up_button(&mut self) {
        self.events.enqueue(Event::PushedVolumeUpButton);
    }

    pub fn be_pushed_volume_down_button(&mut self) {
        self.events.enqueue(Event::PushedVolumeDownButton);
    }

    pub fn handle_event(&mut self, event: Event) {
        match &self.state {
            &State::PowerOn => match event {
                Event::PushedPowerButton => {
                    self.state = State::PowerOff;
                }
                Event::PushedVolumeUpButton => {
                    self.volume += 1;
                }
                Event::PushedVolumeDownButton => {
                    self.volume -= 1;
                }
            },
            &State::PowerOff => match event {
                Event::PushedPowerButton => {
                    self.state = State::PowerOn;
                }
                _ => (),
            },
        }
    }
}

fn push_random_button_of_tv(tv: &mut TV) {
    let mut rng = rand::thread_rng();
    match rng.gen_range(0..4) {
        1 => tv.be_pushed_power_button(),
        2 => tv.be_pushed_volume_up_button(),
        3 => tv.be_pushed_volume_down_button(),
        _ => (),
    };
}

fn main() {
    // println!("Hello, world!");
    // let s = State::PowerOff;
    // println!("show state {:?}", s);
    // let tv = TV::new();
    // println!("show tv {:?}", tv);

    let mut tv = TV::new();
    tv.be_pushed_power_button();
    loop {
        push_random_button_of_tv(&mut tv);
        if let Some(event) = tv.events.dequeue() {
            println!(
                "tv info: {{ state = {:?}, volume = {} }}\ninput_event = {:?}",
                tv.state, tv.volume, event
            );
            tv.handle_event(event);
        }
        thread::sleep(Duration::from_secs(2));
    }
    // Output
    // 23:16:23 > cargo run
    //   Compiling tv-state v0.1.0 (/home/ochi/github.com/ddddddO/work/rust/tv-state)
    //    Finished dev [unoptimized + debuginfo] target(s) in 1.74s
    //     Running `target/debug/tv-state`
    //  tv info: { state = PowerOff, volume = 10 }
    //  input_event = PushedPowerButton
    //  tv info: { state = PowerOn, volume = 10 }
    //  input_event = PushedVolumeDownButton
    //  tv info: { state = PowerOn, volume = 9 }
    //  input_event = PushedVolumeUpButton
    //  tv info: { state = PowerOn, volume = 10 }
    //  input_event = PushedVolumeDownButton
    //  tv info: { state = PowerOn, volume = 9 }
    //  input_event = PushedPowerButton
    //  tv info: { state = PowerOff, volume = 9 }
    //  input_event = PushedVolumeDownButton
    //  tv info: { state = PowerOff, volume = 9 }
    //  input_event = PushedVolumeDownButton
    //  tv info: { state = PowerOff, volume = 9 }
    //  input_event = PushedVolumeUpButton
}
