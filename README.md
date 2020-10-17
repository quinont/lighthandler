# lighthandler

La idea de esto es que te suscribas a una cola y cuando se escriba algo se manda una alerta a una luz Yeelight.


## Variables

estas son las variables:

- IP_LAMP: ip de la lampara.
- IP_MQTT: ip del mosquitto.
- TOPIC: topico donde se suscribe la app

## Ejecutar en docker

de esta forma se ejecuta en docker, primero hay que hacer la imagen:
```
docker build -t lighthandler:1 .
```

despues correr docker run con las variables de entorno adecuadas
```
docker run -e IP_LAMP="XXX.XXX.XXX.XXX" -e IP_MQTT="XXX.XXX.XXX.XXX" -e TOPIC="#" lighthandler:1
```


TODO:
- Mejorar TO-DO.
- Mejorar la documentacion.
- Mejorar codigo.
