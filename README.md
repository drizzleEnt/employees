# employees
http service for employees that provides data about them

POST for Create new employee "/"
PUT for Update employee "/:id"
DELETE for delete employee "/:id"
GET "/company/:id" to get all employees with this company
    "/departament/company_id/departament_name" to get all employees with this department in company


json request:

{
    "name" : "name",
    "surname":"surname",
    "phone": "phone",
    "companyId": company id,
    "passport":{
        "type":"type",
        "number":"number"
    },
    "departament":{
        "name":"name",
        "phone":"phone"
    }
}