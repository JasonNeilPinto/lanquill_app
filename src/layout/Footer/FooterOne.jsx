/* eslint-disable react/prop-types */
import React from "react";
import { Link } from "react-router-dom";
import { FaFacebook, FaGithub, FaDribbble, FaTwitter } from "react-icons/fa";

const FooterOne = ({ footerLight, style, footerGradient }) => {
  return (
    <>
      <footer className="footer-section">
        <div
          className={`footer-top ${footerLight ? "footer-light" : "bg-dark"} ${
            footerGradient ? "bg-gradient" : ""
          }  text-white ptb-120`}
          style={style}
        >
          <div className="container">
            <div className="row justify-content-between">
              <div className="col-md-8 col-lg-4 mb-md-4 mb-lg-0">
                <div className="footer-single-col">
                  <div className="footer-single-col mb-4">
                    <img
                      src="/img/logo-white.png"
                      alt="logo"
                      className="img-fluid logo-white"
                    />
                    <img
                      src="/img/logo-color.png"
                      alt="logo"
                      className="img-fluid logo-color"
                    />
                  </div>
                  <p>
                    Our latest news, articles, and resources, we will sent to
                    your inbox weekly.
                  </p>

                  <form className="newsletter-form position-relative d-block d-lg-flex d-md-flex">
                    <input
                      type="text"
                      className="input-newsletter form-control me-2"
                      placeholder="Enter your email"
                      name="email"
                      required=""
                      autoComplete="off"
                    />
                    <input
                      type="submit"
                      value="Subscribe"
                      data-wait="Please wait..."
                      className="btn btn-primary mt-3 mt-lg-0 mt-md-0"
                    />
                  </form>
                  <div className="ratting-wrap mt-4">
                    <h6 className="mb-0">10/10 Overall rating</h6>
                    <ul className="list-unstyled rating-list list-inline mb-0">
                      <li className="list-inline-item">
                        <i className="fas fa-star text-warning"></i>
                      </li>
                      <li className="list-inline-item">
                        <i className="fas fa-star text-warning"></i>
                      </li>
                      <li className="list-inline-item">
                        <i className="fas fa-star text-warning"></i>
                      </li>
                      <li className="list-inline-item">
                        <i className="fas fa-star text-warning"></i>
                      </li>
                      <li className="list-inline-item">
                        <i className="fas fa-star text-warning"></i>
                      </li>
                    </ul>
                  </div>
                </div>
              </div>
              <div className="col-md-12 col-lg-7 mt-4 mt-md-0 mt-lg-0">
                <div className="row">
                  <div className="col-md-4 col-lg-4 mt-4 mt-md-0 mt-lg-0">
                    <div className="footer-single-col">
                      <h3>Primary Pages</h3>
                      <ul className="list-unstyled footer-nav-list mb-lg-0">
                        <li>
                          <Link to="/" className="text-decoration-none">
                            Home
                          </Link>
                        </li>
                        <li>
                          <Link to="/about-us" className="text-decoration-none">
                            About Us
                          </Link>
                        </li>
                        <li>
                          <Link to="/services" className="text-decoration-none">
                            Services
                          </Link>
                        </li>
                        <li>
                          <Link to="/career" className="text-decoration-none">
                            Career
                          </Link>
                        </li>
                        <li>
                          <Link
                            to="/integrations"
                            className="text-decoration-none"
                          >
                            Integrations
                          </Link>
                        </li>
                        <li>
                          <Link
                            to="/integration-single"
                            className="text-decoration-none"
                          >
                            Integration Single
                          </Link>
                        </li>
                      </ul>
                    </div>
                  </div>
                  <div className="col-md-4 col-lg-4 mt-4 mt-md-0 mt-lg-0">
                    <div className="footer-single-col">
                      <h3>Pages</h3>
                      <ul className="list-unstyled footer-nav-list mb-lg-0">
                        <li>
                          <Link to="/pricing" className="text-decoration-none">
                            Pricing
                          </Link>
                        </li>
                        <li>
                          <Link to="/blogs" className="text-decoration-none">
                            Blog
                          </Link>
                        </li>
                        <li>
                          <Link
                            to="/blog-single"
                            className="text-decoration-none"
                          >
                            Blog Details
                          </Link>
                        </li>
                        <li>
                          <Link
                            to="/contact-us"
                            className="text-decoration-none"
                          >
                            Contact
                          </Link>
                        </li>
                        <li>
                          <Link
                            to="/career-single"
                            className="text-decoration-none"
                          >
                            Career Single
                          </Link>
                        </li>
                        <li>
                          <Link
                            to="/single-service"
                            className="text-decoration-none"
                          >
                            Services Single
                          </Link>
                        </li>
                      </ul>
                    </div>
                  </div>
                  <div className="col-md-4 col-lg-4 mt-4 mt-md-0 mt-lg-0">
                    <div className="footer-single-col">
                      <h3>Template</h3>
                      <ul className="list-unstyled footer-nav-list mb-lg-0">
                        <li>
                          <Link
                            to="/contact-us"
                            className="text-decoration-none"
                          >
                            Contact
                          </Link>
                        </li>
                        <li>
                          <Link
                            to="/help-center"
                            className="text-decoration-none"
                          >
                            Support
                          </Link>
                        </li>
                        <li>
                          <Link
                            to="/help-center-single"
                            className="text-decoration-none"
                          >
                            Support Single
                          </Link>
                        </li>
                        <li>
                          <Link
                            to="/coming-soon"
                            className="text-decoration-none"
                          >
                            Coming Soon
                          </Link>
                        </li>
                        <li>
                          <Link to="/login" className="text-decoration-none">
                            User Login
                          </Link>
                        </li>
                        <li>
                          <Link
                            to="/career-single"
                            className="text-decoration-none"
                          >
                            Career Single
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

        <div
          className={`footer-bottom ${
            footerLight ? "footer-light" : "bg-dark"
          } ${footerGradient ? "bg-gradient" : ""} text-white py-4`}
        >
          <div className="container">
            <div className="row justify-content-between align-items-center">
              <div className="col-md-7 col-lg-7">
                <div className="copyright-text">
                  <p className="mb-lg-0 mb-md-0">
                    &copy; 2021 Quiety Rights Reserved. Designed By{" "}
                    <Link
                      to="/https://themetags.com/"
                      className="text-decoration-none ms-2"
                    >
                      ThemeTags
                    </Link>
                  </p>
                </div>
              </div>
              {/* <div className="col-md-4 col-lg-4">
                <div className="footer-single-col text-start text-lg-end text-md-end">
                  <ul className="list-unstyled list-inline footer-social-list mb-0">
                    <li className="list-inline-item">
                      <Link to="/#">
                        <i className="fab fa-facebook-f"></i>
                      </Link>
                    </li>
                    <li className="list-inline-item">
                      <Link to="/#">
                        <i className="fab fa-instagram"></i>
                      </Link>
                    </li>
                    <li className="list-inline-item">
                      <Link to="/#">
                        <i className="fab fa-dribbble"></i>
                      </Link>
                    </li>
                    <li className="list-inline-item">
                      <Link to="/#">
                        <i className="fab fa-github"></i>
                      </Link>
                    </li>
                  </ul>
                </div>
              </div> */}
              <div className="col-md-4 col-lg-4">
                <div className="footer-single-col text-start text-lg-end text-md-end">
                  <ul className="list-unstyled list-inline footer-social-list mb-0">
                    <li className="list-inline-item">
                      <a href="#">
                        <FaFacebook />
                      </a>
                    </li>
                    <li className="list-inline-item">
                      <a href="#">
                        <FaTwitter />
                      </a>
                    </li>
                    <li className="list-inline-item">
                      <a href="#">
                        <FaDribbble />
                      </a>
                    </li>
                    <li className="list-inline-item">
                      <a href="#">
                        <FaGithub />
                      </a>
                    </li>
                  </ul>
                </div>
              </div>
            </div>
          </div>
        </div>
      </footer>
    </>
  );
};

export default FooterOne;
