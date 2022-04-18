fn main() {

    let mut index: f64 = 0.0;

    while (index >= 0.0) {

        if (index == 0) {
            index = index + 100;
        } else if (index > 50) {
            index = index / 2 - 25;
        } else {
            index = (index / 2) - 1;
        } 

        println!("{}",index);
    }
    println!("{}",index);
}