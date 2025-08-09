/* eslint-disable react/prop-types */
import React, { useEffect, useState } from "react";
import { Link, NavLink, useLocation } from "react-router-dom";
import OffCanvasMenu from "./OffCanvasMenu";
import { HiMenu, HiOutlineX } from "react-icons/hi";
import {
  navServiceLinks,
  navInsightsLinks,
  navProductsLinks,
} from "../../data";
const Navbar = ({
  navDark,
  insurance,
  classOption,
  corporate,
  creativeAgencyOne,
  itCompany,
}) => {
  const [scroll, setScroll] = useState(0);
  const [headerTop, setHeaderTop] = useState(0);

  useEffect(() => {
    const stickyheader = document.querySelector(".main-header");
    setHeaderTop(stickyheader.offsetTop);
    window.addEventListener("scroll", handleScroll);
    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);

  const handleScroll = () => {
    setScroll(window.scrollY);
  };

  // helper: does current location match any child or base paths?
  function useMenuActive(childHrefs = [], basePaths = []) {
    const { pathname, hash } = useLocation();

    const inBase = basePaths.some(
      (p) => pathname === p || pathname.startsWith(`${p}/`)
    );

    const inChildren = childHrefs.some((h) => {
      if (!h) return false;

      // hash links (e.g., "#pramitihr")
      if (h.startsWith("#")) return hash === h;

      // normal routes, allow descendants; ignore child hashes if any
      const [pathOnly] = h.split("#");
      return pathname === pathOnly || pathname.startsWith(`${pathOnly}/`);
    });

    return inBase || inChildren;
  }

  const servicesActive = useMenuActive(
    navServiceLinks.map((l) => l.href),
    ["/services"]
  );

  const resourcesActive = useMenuActive(
    navInsightsLinks.map((l) => l.href),
    ["/resources", "/blog"]
  );

  return (
    <>
      <header
        className={`main-header z-10 ${
          creativeAgencyOne ? "creative_agency_nav " : ""
        }  ${itCompany ? "it_company_nav " : ""}     ${
          corporate ? "header-35 position-absolute top-0 start-0 zindex-9" : ""
        }   ${navDark ? "position-absolute " : ""} w-100 ${classOption} ${
          insurance && "ins-header main-header w-100 z-10 "
        }`}
      >
        <nav
          className={`navbar navbar-expand-xl z-50  ${
            corporate ? "affix" : ""
          } ${navDark ? "navbar-dark " : "navbar-light"} sticky-header ${
            scroll > headerTop ? "affix" : ""
          }`}
        >
          <div className="container d-flex align-items-center justify-content-lg-between position-relative">
            <Link to="/">
              {scroll > headerTop || !navDark || itCompany ? (
                <img
                  width={113}
                  height={36}
                  src="/img/logo-na-white.png"
                  alt="logo"
                  className="img-fluid logo-color"
                />
              ) : (
                <img
                  width={113}
                  height={36}
                  src="/img/logo-na.png"
                  alt="logo"
                  className="img-fluid logo-white"
                />
              )}
            </Link>
            <button
              className="navbar-toggler position-absolute right-0 border-0"
              id="#offcanvasWithBackdrop"
              role="button"
            >
              <span
                data-bs-toggle="offcanvas"
                data-bs-target="#offcanvasWithBackdrop"
                aria-controls="offcanvasWithBackdrop"
              >
                <HiMenu />
              </span>
            </button>
            <div className="clearfix"></div>
            <div className="collapse navbar-collapse justify-content-center">
              <ul className="nav col-12 col-md-auto justify-content-center main-menu">
                <li>
                  <NavLink
                    to="/"
                    className={({ isActive }) =>
                      `nav-link ${isActive ? "active" : ""}`
                    }
                  >
                    Home
                  </NavLink>
                </li>
                <li className="nav-item dropdown">
                  <a
                    className="nav-link dropdown-toggle"
                    href="#"
                    role="button"
                    data-bs-toggle="dropdown"
                    aria-expanded="false"
                  >
                    Products
                  </a>
                  <div className="dropdown-menu border-0 rounded-custom shadow py-0 bg-white">
                    <div className="dropdown-grid rounded-custom width-full">
                      <div className="dropdown-grid-item">
                        <h6 className="drop-heading">Our Products</h6>
                        {navProductsLinks.map((navLink, i) => (
                          <div key={i + 1}>
                            <NavLink
                              to={navLink.href}
                              className={({ isActive }) =>
                                `dropdown-link px-0 ${isActive ? "active" : ""}`
                              }
                              onClick={(e) => {
                                if (window.location.pathname === "/") {
                                  e.preventDefault();

                                  const hash = navLink.href.startsWith("#")
                                    ? navLink.href
                                    : `#${navLink.href.split("#")[1]}`;

                                  // Scroll to section
                                  const section =
                                    document.querySelector("#div-products");
                                  if (section) {
                                    section.scrollIntoView({
                                      behavior: "smooth",
                                    });
                                  }

                                  // Activate the tab
                                  const tabBtn = document.querySelector(
                                    `a[href="${hash}"]`
                                  );
                                  if (tabBtn instanceof HTMLElement)
                                    tabBtn.click();

                                  // Update URL hash
                                  window.history.pushState(null, "", hash);
                                }
                              }}
                            >
                              <span className="me-2">{navLink.icon}</span>
                              <span className="drop-title mb-0">
                                {navLink.title}
                              </span>
                            </NavLink>
                          </div>
                        ))}
                      </div>
                    </div>
                  </div>
                </li>

                <li className="nav-item dropdown">
                  <a
                    className={`nav-link dropdown-toggle ${
                      servicesActive ? "active" : ""
                    }`}
                    href="#"
                    role="button"
                    data-bs-toggle="dropdown"
                    aria-expanded="false"
                  >
                    Services
                  </a>
                  <div className="dropdown-menu border-0 rounded-custom shadow py-0 bg-white">
                    <div className="dropdown-grid rounded-custom width-full">
                      <div className="dropdown-grid-item">
                        <h6 className="drop-heading">Our Services</h6>
                        {navServiceLinks.map((navLink, i) => (
                          <div key={i + 1}>
                            <NavLink
                              to={navLink.href}
                              className={({ isActive }) =>
                                `dropdown-link px-0 ${isActive ? "active" : ""}`
                              }
                            >
                              <span className="me-2">{navLink.icon}</span>
                              <span className="drop-title mb-0">
                                {navLink.title}
                              </span>
                            </NavLink>
                          </div>
                        ))}
                      </div>
                    </div>
                  </div>
                </li>
                <li className="nav-item dropdown">
                  <a
                    className={`nav-link dropdown-toggle ${
                      resourcesActive ? "active" : ""
                    }`}
                    href="#"
                    role="button"
                    data-bs-toggle="dropdown"
                    aria-expanded="false"
                  >
                    Resources
                  </a>
                  <div className="dropdown-menu border-0 rounded-custom shadow py-0 bg-white">
                    <div className="dropdown-grid rounded-custom width-full">
                      <div className="dropdown-grid-item">
                        <h6 className="drop-heading">Resources</h6>
                        {navInsightsLinks.map((navLink, i) => (
                          <div key={i + 1}>
                            <NavLink
                              to={navLink.href}
                              className={({ isActive }) =>
                                `dropdown-link px-0 ${isActive ? "active" : ""}`
                              }
                            >
                              <span className="me-2">{navLink.icon}</span>
                              <span className="drop-title mb-0">
                                {navLink.title}
                              </span>
                            </NavLink>
                          </div>
                        ))}
                      </div>
                    </div>
                  </div>
                </li>
                <li>
                  <NavLink
                    to="/about-us"
                    className={({ isActive }) =>
                      `nav-link ${isActive ? "active" : ""}`
                    }
                  >
                    About
                  </NavLink>
                </li>
              </ul>
            </div>

            <div className="action-btns text-end me-5 me-lg-0 d-none d-md-block d-lg-block">
              <Link
                to="/contact-us"
                className={
                  insurance ? "ins-btn ins-primary-btn" : "btn btn-primary"
                }
              >
                Contact Us
              </Link>
            </div>

            <div
              className="offcanvas offcanvas-end d-xl-none"
              tabIndex="-1"
              id="offcanvasWithBackdrop"
            >
              <div className="offcanvas-header d-flex align-items-center mt-4">
                <Link
                  to="/"
                  className="d-flex align-items-center mb-md-0 text-decoration-none"
                >
                  <img
                    width={121}
                    height={36}
                    src="/img/logo-na-white.png"
                    alt="logo"
                    className="img-fluid ps-2"
                  />
                </Link>
                <button
                  type="button"
                  className="close-btn text-danger"
                  data-bs-dismiss="offcanvas"
                  aria-label="Close"
                >
                  <HiOutlineX />
                </button>
              </div>

              <OffCanvasMenu />
            </div>
          </div>
        </nav>
      </header>
    </>
  );
};

export default Navbar;
