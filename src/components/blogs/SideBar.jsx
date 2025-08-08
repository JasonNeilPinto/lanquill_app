import React from "react";

const SideBar = ({ menu }) => {
  const handleScroll = (id) => {
    const element = document.getElementById(id.replace("#", ""));
    if (element) {
      const headerOffset = 90;
      const elementPosition =
        element.getBoundingClientRect().top + window.pageYOffset;
      const offsetPosition = elementPosition - headerOffset;

      window.scrollTo({
        top: offsetPosition,
        behavior: "smooth",
      });
    }
  };

  return (
    <div
      className="bg-light p-4 sticky-top rounded shadow-sm mt-lg-0"
      style={{ top: "100px", minWidth: "220px" }}
    >
      <h2 className="mb-3">Contents</h2>
      <ul className="list-unstyled mb-0">
        {menu.map((item) => (
          <li key={item.id} className="mb-2">
            <a
              href={item.id}
              className="text-dark text-decoration-none"
              style={{ cursor: "pointer" }}
              onClick={(e) => {
                e.preventDefault();
                handleScroll(item.id);
              }}
            >
              {item.label}
            </a>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default SideBar;
