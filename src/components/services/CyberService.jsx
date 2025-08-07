import React from "react";
import { Link } from "react-router-dom";

const CyberService = () => {
  return (
    <section className="cyber-features pt-100 bg-light">
      <div className="container">
        <div className="row justify-content-center">
          <div className="col-md-8 col-lg-6">
            <div className="section-heading text-center mb-5">
              <h5 className="h6 text-primary">Service</h5>
              <h2>High Quality Trusted Cyber Security solution</h2>
              <p>
                Uniquely promote adaptive quality vectors rather than
                stand-alone e-markets. pontificate alternative architectures
                whereas iterate
              </p>
            </div>
          </div>
        </div>
        <div className="row">
          <div className="col-lg-4">
            <div
              className="cyber-single-service bg-white rounded-custom mb-30"
              style={{
                minHeight: "350px",
              }}
            >
              <div className="feature-icon pb-5 rounded bg-primary-soft text-white mb-4">
                <i className="far fa-chart-line"></i>
              </div>
              <div className="cyber feature-info-wrap">
                <h3 className="h5">Data & Analytics:</h3>
                <p>
                  Transform your data into insights with seamless integration,
                  scalable storage, and smart dashboards.
                </p>
              </div>
              <Link
                to="/single-service"
                className="link-with-icon text-decoration-none"
              >
                Explore More <i className="far fa-arrow-right"></i>
              </Link>
            </div>
          </div>
          <div className="col-lg-4">
            <div
              className="cyber-single-service bg-white rounded-custom mb-30"
              style={{
                minHeight: "350px",
              }}
            >
              <div className="feature-icon pb-5 rounded bg-primary-soft text-white mb-4">
                <i className="far fa-shield-alt"></i>
              </div>
              <div className="cyber feature-info-wrap">
                <h3 className="h5">Cyber Security:</h3>
                <p>
                  Comprehensive cybersecurity with compliance, threat
                  intelligence, and DevSecOps for secure, scalable
                  infrastructure.
                </p>
              </div>
              <Link
                to="/single-service"
                className="link-with-icon text-decoration-none"
              >
                Explore More <i className="far fa-arrow-right"></i>
              </Link>
            </div>
          </div>
          <div className="col-lg-4">
            <div
              className="cyber-single-service bg-white rounded-custom mb-30"
              style={{
                minHeight: "350px",
              }}
            >
              <div className="feature-icon pb-5 rounded bg-primary-soft text-white mb-4">
                <i className="far fa-cloud"></i>
              </div>
              <div className="cyber feature-info-wrap">
                <h3 className="h5">Cloud Services: </h3>
                <p>
                  Modernize applications, streamline cloud strategies, and
                  optimize costs with tailored multi-cloud, hybrid cloud, and
                  cloud migration solutions.
                </p>
              </div>
              <Link
                to="/single-service"
                className="link-with-icon text-decoration-none"
              >
                Explore More <i className="far fa-arrow-right"></i>
              </Link>
            </div>
          </div>
          <div className="col-lg-4">
            <div
              className="cyber-single-service bg-white rounded-custom mb-30"
              style={{
                minHeight: "350px",
              }}
            >
              <div className="feature-icon pb-5 rounded bg-primary-soft text-white mb-4">
                <i className="far fa-brain"></i>
              </div>
              <div className="cyber feature-info-wrap">
                <h3 className="h5">Applied AI:</h3>
                <p>
                  Leverage generative AI and machine learning to streamline
                  operations and drive data-driven decisions.
                </p>
              </div>
              <Link
                to="/single-service"
                className="link-with-icon text-decoration-none"
              >
                Explore More <i className="far fa-arrow-right"></i>
              </Link>
            </div>
          </div>
          <div className="col-lg-4">
            <div
              className="cyber-single-service bg-white rounded-custom mb-30"
              style={{
                minHeight: "350px",
              }}
            >
              <div className="feature-icon pb-5 rounded bg-primary-soft text-white mb-4">
                <i className="far fa-bug"></i>
              </div>
              <div className="cyber feature-info-wrap">
                <h3 className="h5">Security testing: </h3>
                <p>
                  Detect vulnerabilities early with security testing services
                  like penetration testing, compliance checks, red teaming and
                  risk assessment—built to protect and scale.
                </p>
              </div>
              <Link
                to="/single-service"
                className="link-with-icon text-decoration-none"
              >
                Explore More <i className="far fa-arrow-right"></i>
              </Link>
            </div>
          </div>
          <div className="col-lg-4">
            <div
              className="cyber-single-service bg-white rounded-custom mb-30"
              style={{
                minHeight: "350px",
              }}
            >
              <div className="feature-icon pb-5 rounded bg-primary-soft text-white mb-4">
                <i className="far fa-shield-alt"></i>
              </div>
              <div className="cyber feature-info-wrap">
                <h3 className="h5">Managed Services</h3>
                <p>
                  Randomised words which don&apos;t look even passage of Lorem
                  Ipsum. You need to be Lorem Ipsum randomised even .
                </p>
              </div>
              <Link
                to="/single-service"
                className="link-with-icon text-decoration-none"
              >
                Explore More <i className="far fa-arrow-right"></i>
              </Link>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default CyberService;
