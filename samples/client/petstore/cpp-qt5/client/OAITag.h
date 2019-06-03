/**
 * OpenAPI Petstore
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/*
 * OAITag.h
 *
 * A tag for a pet
 */

#ifndef OAITag_H
#define OAITag_H

#include <QJsonObject>


#include <QString>

#include "OAIObject.h"
#include "OAIEnum.h"

namespace OpenAPI {

class OAITag: public OAIObject {
public:
    OAITag();
    OAITag(QString json);
    ~OAITag() override;

    QString asJson () const override;
    QJsonObject asJsonObject() const override;
    void fromJsonObject(QJsonObject json) override;
    void fromJson(QString jsonString) override;

    
    qint64 getId() const;
    void setId(const qint64 &id);

    
    QString getName() const;
    void setName(const QString &name);

    
    
    virtual bool isSet() const override;
    virtual bool isValid() const override;

private:
    void init();
    
    qint64 id;
    bool m_id_isSet;
    bool m_id_isValid;
    
    QString name;
    bool m_name_isSet;
    bool m_name_isValid;
    
    };

}

#endif // OAITag_H
