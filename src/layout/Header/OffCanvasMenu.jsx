import React from "react";
import { Link } from "react-router-dom";
import {
  navServiceLinks,
  navInsightsLinks,
  navProductsLinks,
} from "../../data";

const OffCanvasMenu = () => {
  return (
    <div className="offcanvas-body">
      <ul className="nav col-12 col-md-auto justify-content-center main-menu">
        {/* Home */}
        <li data-bs-dismiss="offcanvas" aria-label="Close">
          <Link to="/" className="nav-link">
            Home
          </Link>
        </li>

        {/* Products dropdown */}
        <li className="nav-item dropdown">
          <span
            className="nav-link dropdown-toggle d-flex justify-content-between"
            role="button"
            data-bs-toggle="dropdown"
            aria-expanded="false"
          >
            Products
          </span>
          <div className="dropdown-menu border-0 rounded-custom shadow py-0 bg-white">
            <div className="dropdown-grid rounded-custom width-full">
              <div className="dropdown-grid-item">
                <h6 className="drop-heading">Our Products</h6>
                {navProductsLinks.map((navLink, i) => (
                  <Link
                    key={i}
                    to={navLink.href}
                    className="dropdown-link px-0 d-flex align-items-center"
                    data-bs-dismiss="offcanvas"
                  >
                    <span className="me-2">{navLink.icon}</span>
                    <span className="drop-title mb-0">{navLink.title}</span>
                  </Link>
                ))}
              </div>
            </div>
          </div>
        </li>

        {/* Services dropdown */}
        <li className="nav-item dropdown">
          <span
            className="nav-link dropdown-toggle d-flex justify-content-between"
            role="button"
            data-bs-toggle="dropdown"
            aria-expanded="false"
          >
            Services
          </span>
          <div className="dropdown-menu border-0 rounded-custom shadow py-0 bg-white">
            <div className="dropdown-grid rounded-custom width-full">
              <div className="dropdown-grid-item">
                <h6 className="drop-heading">Our Services</h6>
                {navServiceLinks.map((navLink, i) => (
                  <Link
                    key={i}
                    to={navLink.href}
                    className="dropdown-link px-0 d-flex align-items-center"
                    data-bs-dismiss="offcanvas"
                  >
                    <span className="me-2">{navLink.icon}</span>
                    <span className="drop-title mb-0">{navLink.title}</span>
                  </Link>
                ))}
              </div>
            </div>
          </div>
        </li>

        {/* Resources dropdown */}
        <li className="nav-item dropdown">
          <span
            className="nav-link dropdown-toggle d-flex justify-content-between"
            role="button"
            data-bs-toggle="dropdown"
            aria-expanded="false"
          >
            Resources
          </span>
          <div className="dropdown-menu border-0 rounded-custom shadow py-0 bg-white">
            <div className="dropdown-grid rounded-custom width-full">
              <div className="dropdown-grid-item">
                <h6 className="drop-heading">Resources</h6>
                {navInsightsLinks.map((navLink, i) => (
                  <Link
                    key={i}
                    to={navLink.href}
                    className="dropdown-link px-0 d-flex align-items-center"
                    data-bs-dismiss="offcanvas"
                  >
                    <span className="me-2">{navLink.icon}</span>
                    <span className="drop-title mb-0">{navLink.title}</span>
                  </Link>
                ))}
              </div>
            </div>
          </div>
        </li>

        {/* About */}
        <li data-bs-dismiss="offcanvas" aria-label="Close">
          <Link to="/about-us" className="nav-link">
            About
          </Link>
        </li>
      </ul>

      {/* Action buttons */}
      <div className="action-btns mt-4 ps-3">
        <Link
          to="/login"
          className="btn btn-outline-primary text-decoration-none me-2"
          data-bs-dismiss="offcanvas"
          aria-label="Close"
        >
          Sign In
        </Link>
        <Link
          to="/request-demo"
          className="btn btn-primary"
          data-bs-dismiss="offcanvas"
          aria-label="Close"
        >
          Get Started
        </Link>
      </div>
    </div>
  );
};

export default OffCanvasMenu;
