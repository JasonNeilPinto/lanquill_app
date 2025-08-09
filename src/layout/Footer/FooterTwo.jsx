import React from "react";
import { Link } from "react-router-dom";

const FooterTwo = () => {
  return (
    <footer className="cyber-footer bg-gradient">
      <div className="cyber-footer-top ptb-80">
        <div className="container">
          <div className="row">
            <div className="col-md-8 col-lg-4 mb-md-4 mb-lg-0">
              <div className="">
                <div className="footer-single-col mb-4">
                  <img
                    src="/img/logo-na.png"
                    alt="logo"
                    className="img-fluid logo-white"
                  />
                </div>
                <p className="text-white">
                  Global IT Consulting and Product Development Company
                  specializing in new-age technologies.
                </p>
                <ul className="list-unstyled list-inline cyber-footer-social-list mb-0">
                  <li className="list-inline-item">
                    <a href="/#">
                      <i className="fab fa-facebook-f"></i>
                    </a>
                  </li>
                  <li className="list-inline-item">
                    <a href="/#">
                      <i className="fab fa-instagram"></i>
                    </a>
                  </li>
                  <li className="list-inline-item">
                    <a href="/#">
                      <i className="fab fa-youtube"></i>
                    </a>
                  </li>
                  <li className="list-inline-item">
                    <a href="/#">
                      <i className="fab fa-linkedin"></i>
                    </a>
                  </li>
                </ul>
              </div>
            </div>
            <div className="col-md-12 col-lg-8 mt-4 mt-md-0 mt-lg-0">
              <div className="row">
                <div className="col-lg-3 col-md-6">
                  <div className="">
                    <h3 className="h5 mb-4 text-white">Services</h3>
                    <ul className="list-unstyled footer-nav-list mb-lg-0">
                      <li>
                        <Link to="/" className="text-decoration-none">
                          Data & Analytics
                        </Link>
                      </li>
                      <li>
                        <Link to="/about-us" className="text-decoration-none">
                          Cyber Security
                        </Link>
                      </li>
                      <li>
                        <Link to="/services" className="text-decoration-none">
                          Cloud Services
                        </Link>
                      </li>
                      <li>
                        <Link to="/career" className="text-decoration-none">
                          Applied AI
                        </Link>
                      </li>
                      <li>
                        <Link
                          to="/integrations"
                          className="text-decoration-none"
                        >
                          Security testing
                        </Link>
                      </li>
                      <li>
                        <Link
                          to="/integration-single"
                          className="text-decoration-none"
                        >
                          Managed Services
                        </Link>
                      </li>
                    </ul>
                  </div>
                </div>
                <div className="col-lg-3 col-md-6">
                  <div className="">
                    <h3 className="h5 mb-4 text-white">Products</h3>
                    <ul className="list-unstyled footer-nav-list mb-lg-0">
                      <li>
                        <Link to="/contact-us" className="text-decoration-none">
                          PramitiHR
                        </Link>
                      </li>
                      <li>
                        <Link to="/about-us" className="text-decoration-none">
                          Lanquill
                        </Link>
                      </li>
                      <li>
                        <Link to="/services" className="text-decoration-none">
                          Gen AI Sandbox
                        </Link>
                      </li>
                      <li>
                        <Link to="/career" className="text-decoration-none">
                          Reverse Auction e-commerce
                        </Link>
                      </li>
                      <li>
                        <Link
                          to="/integrations"
                          className="text-decoration-none"
                        >
                          OnDemand Entertainment
                        </Link>
                      </li>
                    </ul>
                  </div>
                </div>
                <div className="col-lg-3 col-md-6">
                  <div className="">
                    <h3 className="h5 mb-4 text-white">Quick Links</h3>
                    <ul className="list-unstyled footer-nav-list mb-lg-0">
                      <li>
                        <Link to="/" className="text-decoration-none">
                          Contact Us
                        </Link>
                      </li>
                      <li>
                        <Link to="/about-us" className="text-decoration-none">
                          About Us
                        </Link>
                      </li>
                      <li>
                        <Link to="/" className="text-decoration-none">
                          Resources
                        </Link>
                      </li>
                    </ul>
                  </div>
                </div>
                <div className="col-lg-3 col-md-6">
                  <div className="">
                    <h3 className="h5 mb-4 text-white">Contact Info</h3>
                    <ul className="list-unstyled footer-nav-list mb-lg-0">
                      <li>
                        <Link to="/contact-us" className="text-decoration-none">
                          +91 9620555571
                        </Link>
                      </li>
                      <li>
                        <Link to="/contact-us" className="text-decoration-none">
                          support@netanalytiks.com
                        </Link>
                      </li>
                      <li>
                        <Link to="/" className="text-decoration-none">
                          91Springboard, MG Road, Bangalore, India
                        </Link>
                      </li>
                    </ul>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div className="cyber-footer-bottom">
        <div className="container">
          <p className="mb-0 py-4 text-center">
            Copyright ©2025 NetAnalytiks Technologies Ltd. All rights reserved.
          </p>
        </div>
      </div>
    </footer>
  );
};

export default FooterTwo;
