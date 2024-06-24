package main

type Vulnerability struct {
    Lesson       string `json:"lesson"`
    Description  string `json:"description"`
    Remediation  string `json:"remediation"`
}
var vulnerabilities = []Vulnerability{
    {
        Lesson:      "LESSON #1: Event Injection",
        Description: "La aplicación es vulnerable a la inyección de código a través de deserialización insegura en las solicitudes que terminan siendo manejadas por una función Lambda.",
        Remediation: "Use serialización segura y valide las entradas antes de procesarlas.",
    },
    {
        Lesson:      "LESSON #2: Broken Authentication",
        Description: "La aplicación no verifica la firma del JWT en el encabezado de autorización, permitiendo que los usuarios se hagan pasar por otros.",
        Remediation: "Verifique la firma del JWT en cada solicitud para asegurar la autenticación adecuada.",
    },
    {
        Lesson:      "LESSON #3: Sensitive Data Exposure",
        Description: "Un atacante puede invocar una función de administración que revela recibos del depósito S3 a través de una URL firmada.",
        Remediation: "Restrinja el acceso a las funciones administrativas y asegure los datos sensibles en S3.",
    },
    {
        Lesson:      "LESSON #4: Denial of Service (DoS)",
        Description: "La aplicación permite solo 10 invocaciones simultáneas del proceso de facturación, lo que puede ser explotado para causar una Denegación de Servicio.",
        Remediation: "Implemente mecanismos de protección contra DoS y límite la tasa de solicitudes.",
    },
}
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Vulnerability struct {
    Lesson       string `json:"lesson"`
    Description  string `json:"description"`
    Remediation  string `json:"remediation"`
}

var vulnerabilities = []Vulnerability{
    {
        Lesson:      "LESSON #1: Event Injection",
        Description: "La aplicación es vulnerable a la inyección de código a través de deserialización insegura en las solicitudes que terminan siendo manejadas por una función Lambda.",
        Remediation: "Use serialización segura y valide las entradas antes de procesarlas.",
    },
    {
        Lesson:      "LESSON #2: Broken Authentication",
        Description: "La aplicación no verifica la firma del JWT en el encabezado de autorización, permitiendo que los usuarios se hagan pasar por otros.",
        Remediation: "Verifique la firma del JWT en cada solicitud para asegurar la autenticación adecuada.",
    },
    {
        Lesson:      "LESSON #3: Sensitive Data Exposure",
        Description: "Un atacante puede invocar una función de administración que revela recibos del depósito S3 a través de una URL firmada.",
        Remediation: "Restrinja el acceso a las funciones administrativas y asegure los datos sensibles en S3.",
    },
    {
        Lesson:      "LESSON #4: Denial of Service (DoS)",
        Description: "La aplicación permite solo 10 invocaciones simultáneas del proceso de facturación, lo que puede ser explotado para causar una Denegación de Servicio.",
        Remediation: "Implemente mecanismos de protección contra DoS y límite la tasa de solicitudes.",
    },
}

func main() {
    r := gin.Default()

    r.GET("/vulnerabilities", func(c *gin.Context) {
        c.JSON(http.StatusOK, vulnerabilities)
    })

    r.Run(":8080") // Corre la aplicación en el puerto 8080
}
