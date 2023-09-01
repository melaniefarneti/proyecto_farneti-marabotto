import React from "react";

function Footer() {
  return (
    <footer className="page-footer">
      <div className="container">
        <div className="row">
          <div className="col l6 s12">
            <h5 className="white-text">Arquitectura de Software 2023</h5>
            <p className="grey-text text-lighten-4">Farneti Melanie, Marabotto Tiziano</p>
          </div>
          <div className="col l4 offset-l2 s12">
            <h5 className="white-text">Contacto</h5>
            <ul>
              <li><a className="grey-text text-lighten-3" href="#!">(Melanie) 2113289@ucc.edu.ar</a></li>
              <li><a className="grey-text text-lighten-3" href="#!">(Tiziano)  2112554@ucc.edu.ar</a></li>
            </ul>
          </div>
        </div>
      </div>
      <div className="footer-copyright">
        <div className="container">
          © {new Date().getFullYear()} HiHotels! 
        </div>
      </div>
    </footer>
  );
}

export default Footer;
