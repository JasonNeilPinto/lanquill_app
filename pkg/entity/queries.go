package entity

var updateHigherEntityInfo = `UPDATE 
        sententia.entity
    SET
        Entity_Name = ?, Logo_Path = ?, logo_enable = ?
    WHERE
        Entity_ID = ?`

var updateLowerEntityInfo = `UPDATE 
        sententia.entity_division
    SET
        Entity_Div_Name = ?
    WHERE
        Entity_Div_ID = ?`

var deleteHigherEntity = `UPDATE 
        sententia.entity
    SET
        Entity_Status = "Deleted", Entity_Name = CONCAT(Entity_Name,"_Deleted"), Date_Modified = CURRENT_TIMESTAMP()
    WHERE
        Entity_ID = ? OR Parent_Entity_ID = ?`

var deleteLowerEntityWithParent = `UPDATE 
        sententia.entity_division
    SET
        Entity_Status = "Deleted", Entity_Div_Name = CONCAT(Entity_Div_Name,"_Deleted"), Date_Modified = CURRENT_TIMESTAMP()
    WHERE
        Entity_ID = ?`

var deleteLowerEntity = `UPDATE 
        sententia.entity_division
    SET
        Entity_Status = "Deleted", Entity_Div_Name = CONCAT(Entity_Div_Name,"_Deleted"), Date_Modified = CURRENT_TIMESTAMP()
    WHERE
        Entity_Div_ID = ? OR Parent_Entity_Div_ID = ?`

var deleteEntityUsers = `UPDATE 
        sententia.user
    SET
        User_Status = "Deleted", User_Email = CONCAT(User_Email,"_Deleted"), Last_Modified = CURRENT_TIMESTAMP()
    WHERE
        %s = ?`

var getEntityTypeId = `SELECT Entity_Type_ID FROM sententia.entity_type WHERE Entity_Type_Name = ? AND Parent_Entity_Type_ID != 1`

var getEntityDivisionTypeId = `SELECT Entity_Div_Type_ID FROM sententia.entity_division_type WHERE Entity_Div_Type_Name = ? AND Entity_Type_ID = ?`

var getEntityTypeIdDirect = `SELECT Entity_Type_ID FROM sententia.entity_type WHERE Entity_Type_Name = ? AND Parent_Entity_Type_ID = 1`

var getParentEntityTypeId = `SELECT Entity_Type_ID FROM sententia.entity WHERE Entity_ID = ? AND Entity_Status != 'Deleted'`

var getParentEntityId = `SELECT Parent_Entity_ID FROM sententia.entity WHERE Entity_ID = ? AND Entity_Status != 'Deleted'`

var getEntityName = `SELECT Entity_Name FROM sententia.entity WHERE Entity_Name = ? AND Parent_Entity_ID = ? AND Entity_Status != 'Deleted'`

var getParentEntityLicenseDetails = `SELECT el.Number_of_Licences, el.Licence_Type_ID, el.Licence_Consumed
FROM sententia.entity e 
JOIN sententia.entity_licence el ON el.Entity_ID = e.Parent_Entity_ID
WHERE e.Entity_ID = ? AND e.Entity_Status != 'Deleted'`

var insertEntityDetails = `INSERT INTO sententia.entity (Entity_Type_ID, Parent_Entity_ID, Parent_Entity_Type_ID, Entity_Name, Entity_Contact_Name, Entity_Contact_Email, Entity_Contact_Mobile, Date_Created, Date_Modified, Logo_Path) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

var insertEntityLicenseDetails = `INSERT INTO sententia.entity_licence (Entity_ID, Licence_Type_ID, Number_of_Licences, Date_Purchased, Valid_Till) VALUES (?, ?, ?, ?, ?)`

var updateEntityLicense = `UPDATE sententia.entity_licence SET Licence_Consumed = Licence_Consumed + %d where Licence_Type_ID = ? and Entity_ID = ?`

var getExistingEntityName = `SELECT Entity_Name FROM sententia.entity where Entity_Name = ? AND Entity_Status != 'Deleted'`

var updateParentEntityId = `UPDATE sententia.entity SET Parent_Entity_ID = ? WHERE Entity_ID = ?`

var getEntityDivisionName = `SELECT Entity_Div_Name FROM sententia.entity_division WHERE Entity_Div_Name = ? AND Entity_ID = ? AND Parent_Entity_Div_ID = ? AND Entity_Status != 'Deleted'`

var insertEntityDivisionDetails = `INSERT INTO sententia.entity_division (Entity_ID, Entity_Div_Name,Parent_Entity_Div_ID, Entity_Div_Type_ID, Entity_Type_ID, Entity_Div_Type_Name, Entity_Contact_Name, Entity_Contact_Email, Entity_Contact_Mobile) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

var getParentEntityDivisionName = `SELECT Entity_Div_Name FROM sententia.entity_division WHERE Entity_Div_Name = ? AND Entity_ID = ? AND Entity_Div_ID = Parent_Entity_Div_ID AND Entity_Status != 'Deleted'`

var updateParentEntityDivId = `UPDATE sententia.entity_division SET Parent_Entity_Div_ID = ? WHERE Entity_Div_ID = ?`

var updateLogoPath = `UPDATE sententia.entity SET Logo_Path = ?, logo_enable = ? WHERE Entity_ID = ?`
