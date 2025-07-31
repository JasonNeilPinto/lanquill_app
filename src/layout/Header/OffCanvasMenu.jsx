import React from "react";
import { Link } from "react-router-dom";
import { navCompanyLinks, navCompanyPage, offcanvasMenuData } from "../../data";

const OffCanvasMenu = () => {
  return (
    <div className="offcanvas-body">
      <ul className="nav col-12 col-md-auto justify-content-center main-menu">
        {/* Home dropdown */}
        <li className="nav-item dropdown">
          <span
            className="nav-link dropdown-toggle d-flex justify-content-between"
            role="button"
            data-bs-toggle="dropdown"
            aria-expanded="false"
          >
            Home
          </span>
          <div className="dropdown-menu border-0 rounded-custom shadow py-0 bg-white">
            <div className="dropdown-grid rounded-custom width-half">
              <div className="dropdown-grid-item">
                <h6 className="drop-heading">Different Home</h6>
                {offcanvasMenuData.map((navH, i) => (
                  <Link
                    key={i}
                    to={navH.href}
                    className="dropdown-link d-flex align-items-center"
                    data-bs-dismiss="offcanvas"
                  >
                    <span className="demo-list bg-primary rounded text-white fw-bold me-2">
                      {i + 1}
                    </span>
                    <span className="dropdown-info mb-0">
                      <span className="drop-title">{navH.title}</span>
                      <span>{navH.info}</span>
                    </span>
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

        {/* Services */}
        <li data-bs-dismiss="offcanvas" aria-label="Close">
          <Link to="/services" className="nav-link">
            Services
          </Link>
        </li>

        {/* Pricing */}
        <li data-bs-dismiss="offcanvas" aria-label="Close">
          <Link to="/pricing" className="nav-link">
            Pricing
          </Link>
        </li>

        {/* Company dropdown */}
        <li className="nav-item dropdown">
          <span
            className="nav-link dropdown-toggle d-flex justify-content-between"
            role="button"
            data-bs-toggle="dropdown"
            aria-expanded="false"
          >
            Company
          </span>
          <div className="dropdown-menu border-0 rounded-custom shadow py-0 bg-white">
            <div className="dropdown-grid rounded-custom width-full">
              {/* Useful Links */}
              <div className="dropdown-grid-item">
                <h6 className="drop-heading">Useful Links</h6>
                {navCompanyLinks.map((navLink, i) => (
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

              {/* Utility Pages */}
              <div className="dropdown-grid-item radius-right-side bg-light">
                <h6 className="drop-heading">Utility Pages</h6>
                {navCompanyPage.map((navPage, i) => (
                  <Link
                    key={i}
                    to={navPage.href}
                    className="dropdown-link d-flex align-items-center"
                    data-bs-dismiss="offcanvas"
                  >
                    <span className="me-2">{navPage.icon}</span>
                    <span className="drop-title mb-0">{navPage.title}</span>
                  </Link>
                ))}
              </div>
            </div>
          </div>
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
