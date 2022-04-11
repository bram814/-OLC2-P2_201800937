fn main() {
   let numero = 2;
    
    match numero {
      "1"   => 
          println!("{}",1),
      
      //esto se conoce como brazo
      6 | 7 => { 
          println!("{}",2);
      }
      _ => { 
        println!("Resto de casos");
      }
    }
}
