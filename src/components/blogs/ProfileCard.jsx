import React from "react";
import { Link } from "react-router-dom";

const ProfileCard = ({ name, role, description, logos }) => {
  return (
    <>
      <div className="author-wrap text-center bg-light p-5  rounded-custom mt-5 mt-lg-0">
        <img
          src="/img/team/team-2.jpg"
          alt="author"
          width="120"
          className="img-fluid shadow-sm rounded-circle"
        />
        <div className="author-info my-4">
          <h5 className="mb-0">{name}</h5>
          <span className="small">{role}</span>
        </div>
        <p>{description}</p>
        <ul className="list-unstyled author-social-list list-inline mt-3 mb-0">
          {logos?.map((logo, index) => (
            <li key={index} className="list-inline-item">
              <Link to={logo.link}>
                <i className={logo.icon}></i>
              </Link>
            </li>
          ))}
        </ul>
      </div>
    </>
  );
};

export default ProfileCard;
